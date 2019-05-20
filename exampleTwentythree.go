package main

import "fmt"

func createChannel(msg int) {
	coffee := make(chan int)
	go func() {coffee <- msg}()

	cupNumber := <-coffee
	fmt.Println("Cup Number", cupNumber)

}

func Twentythree() {
	for i := 0; i < 8; i ++ {
		createChannel(i)
	}
}






// create a goroutine channel
// messages := make(chan string)
// // send the channel a message
// go func() { messages <- "this is a message being sent to the message channel"}()
// // collect the message and log it
// msg := <-messages
// fmt.Println(msg)