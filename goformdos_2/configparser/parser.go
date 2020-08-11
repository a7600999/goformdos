package configparser

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// ParseConfig - Parse the Config file to destination map[string]string
// The config must be in the right format
// e.g. Host:google.com
func ParseConfig(dest *map[string]string, filename string, wg *sync.WaitGroup) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ERROR: opening file: %s\n", filename)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalf("ERROR: closing file: %s\n", filename)
		}
	}()

	s := bufio.NewScanner(f)
	var stripped []string
	var tempRef map[string]string
	for s.Scan() {
		stripped = strings.Split(s.Text(), ":")         // split line on :
		stripped[0] = strings.Trim(stripped[0], "\r\n") // strip carriage return and newline
		stripped[1] = strings.Trim(stripped[1], "\r\n")
		tempRef = *dest
		tempRef[stripped[0]] = stripped[1]
		stripped = nil // clear the slice
		tempRef = nil  // Cleart the map
	}

	if err = s.Err(); err != nil {
		log.Fatalf("%s\n", err)
	}

	wg.Done()
}

// ParseJSON --
func ParseJSON(dest *map[string]interface{}, filename string, wg *sync.WaitGroup) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileData, &dest)
	if err != nil {
		log.Fatal(err)
	}

	wg.Done()
}
