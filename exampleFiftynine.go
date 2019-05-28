package main

import (
	"fmt"
	"os"
)

func FiftyNine() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]


	args := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(args)
}