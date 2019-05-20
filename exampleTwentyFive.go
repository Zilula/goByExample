package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	//syncing excution across go routines or workers
	// feels similiar to setTimeout in JS
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func TwentyFive() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

}