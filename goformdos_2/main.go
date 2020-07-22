// Author: PiereLucas

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

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

func validateURL(s string, sync chan<- struct{}) {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		log.Fatalf("%s is invalid url\n", s)
	}
	close(sync)
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
}

func appendToStruct(mode string, syncChannel chan<- string, input map[string]string) {
	for k, v := range input {
		if mode == "FORMS" {
			info.AddForm(k, v)
		} else if mode == "HEADERS" {
			info.AddHeader(k, v)
		} else {
			log.Fatalln("error in func appendToStruct - wrong mode! [FORMS / HEADERS]")
		}
	}

	syncChannel <- ""
}

func main() {

	urlsync := make(chan struct{})
	sync := make(chan string)

	var parseerr string
	var forms = make(map[string]string)
	var headers = make(map[string]string)

	go validateURL(*flagURL, urlsync)
	go configparser.Parse(&forms, *flagForms, sync)
	go configparser.Parse(&headers, *flagHeaders, sync)

	<-urlsync         // wait for validateURL()
	parseerr = <-sync // wait for file append
	if parseerr != "" {
		log.Fatalln(parseerr)
	}

	parseerr = <-sync // wait for file append
	if parseerr != "" {
		log.Fatalln(parseerr)
	}

	info.AddWebaddress(*flagURL) // add url to struct

	go appendToStruct("FORMS", sync, forms)     // Intialize forms
	go appendToStruct("HEADERS", sync, headers) // Intialize headers
	<-sync                                      // wait for appendToStruct
	<-sync                                      // ""

	err := dos.Run(*flagThreads, *flagTime, info)
	if err != nil {
		log.Fatalln(err)
	}
}
