package main

import "os"

func FourtyOne() {
	panic("Some problem")

	_, err := os.Create("tmp/file")
	if err != nil {
		panic(err)
	}
}