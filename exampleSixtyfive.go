package main

import (
	"fmt"
	"os"
)

func SixtyFive() {
	defer fmt.Println("!")


	os.Exit(3)
}