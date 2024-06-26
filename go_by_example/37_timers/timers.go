package main

import (
	"fmt"
	"time"
)

// we often want to execute Go code at some point in the future, or repeatedly at some interval
// Go's built-int timer and ticker features make both of these tasks easy
// we'll look first at times and then at tickers

func main() {
	// timers represent a single event in the future
	// you tell the timer how long you want to wait
	// and it porovides a channel that wil be notified at that time
	// this timer will wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// the <-timer1.C blocks on the timer's channel C until it
	// sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("timer 1 fired")

	// if you just wanted to wait, you could have used time.Sleep
	// one reason a timer may be useful is that you can cancel the timer before it fires
	// here's and example of that
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	// give the timer 2 enough time to fire
	// if it ever was going to
	// to show it is in fact stopped
	time.Sleep(2 * time.Second)
}

// the first time will fire ~2s after we start the program
// but the second should be stopped before it has a chance to fire
