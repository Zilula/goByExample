package main

import "fmt"

func Twentyfour() {
	coffeeCounterChannel := make(chan string, 2)
	go func() {
		fmt.Println(<-coffeeCounterChannel, "listener 3")
	}()
	go func() {
		coffeeCounterChannel <- "First message -- buffered"
		coffeeCounterChannel <- "Second message -- channel"
		coffeeCounterChannel <- "Three message -- channel"
		}()
		
		
	fmt.Println(<-coffeeCounterChannel, "listener 1")
	fmt.Println(<-coffeeCounterChannel, "listener 2")
	fmt.Scanln()
}
