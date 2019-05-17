package main

import "fmt"

func counter(from string) {
	for i := 0; i < 8; i++ {
		fmt.Println(from, ":", i*123)
	}
}

func Twentytwo() {
	counter("direct")

	go counter("goRoutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}