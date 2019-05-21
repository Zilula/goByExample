package main

import "fmt"

func ThirtyOne() {
	queue := make(chan string, 2)

	queue <- "One"
	queue <- "Two"
	close(queue)
	for element := range queue {
		fmt.Println(element)
	}
}