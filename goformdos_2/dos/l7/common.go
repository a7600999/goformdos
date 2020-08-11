// Author: PiereLucas

package l7

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

// random block function
func buildblock(size int) string {
	var block []rune
	for i := 0; i < size; i++ {
		block = append(block, rune(rand.Intn(25)+65))
	}
	return string(block)
}

// validateThreads - warns if more threads starting as cpu cores available
func validateThreads(nthreads int, done chan<- struct{}) {
	nCores := runtime.NumCPU() // runtime.NumCPU() returns the number of available CPU Cores
	if nCores < nthreads {
		log.Printf("WARNING: more routines: %d than cores: %d\n", nthreads, nCores)
		time.Sleep(3 * time.Second)
	}
	close(done)
}
