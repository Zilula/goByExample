package main

import (
	"fmt"
)

// pass this function a channel and string
// passes the channel a message string
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// function takes two channels
// gets the message from 1 channel
// passes that message to another channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// create two channels
// pass ping a channel and a message
// pass the other function both channels
func TwentySix() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed the message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}