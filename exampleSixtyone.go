package main

import (
	"os"
	"strings"
	"fmt"
)

func SixtyOne() {
	os.Setenv("FOO", "1")

	fmt.Println("FOO: ", os.Getenv("FOOD"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}