package main

import (
	"time"
	"fmt"
)

func TwentySeven() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "One"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Two"
	}()

	for i:= 0; i < 2; i ++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received Message: ", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received Message: ", msg2)
		}
	}
}