package main

import (
	"fmt"
	"time"
)

func TwentyEight() {
	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "result 1"
	}()

	select {
	case res := <-chan1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout func 1 ")
	}


	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "result 2"
	}()

	select {
	case res := <-chan2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout func 2 ")
	}

}