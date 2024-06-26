package main

import (
	"fmt"
	"time"
)

// we can use channels to synchronize execution across goroutines
// here's an example of using a blocking receive to wait for a goroutine to finish
// when waiting for multiple goroutines to finish you may prefer to use a WaitGroup

// this is the function we'll run in a goroutine
// the done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Printf("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify that we're done
	done <- true
}

func main() {
	// start a worker goroutine
	// giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// block until we receive a notification from the work on the channel
	<-done
}

// if you removed the <- done line from this program
// the program would exit before the worker event started
