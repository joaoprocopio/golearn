package main

import (
	"fmt"
	"time"
)

// a goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// suppose we have a function call f(s)
	// here's how we'd call that in the usual way
	// running it synchronously
	f("direct")

	// to invoke this function in a goroutine
	// this new goroutine wil execute concurrently with the calling one
	go f("goroutine")

	// you can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// our two function calls running asynchronously in separate goroutines now
	// wait for them to finish (for a more robust approach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}

// when we run this program, we see the output of the blocking call first
// then the output of the two goroutines
// the goroutines output may be interleaved
// because goroutines are being run concurrently by the Go runtime
