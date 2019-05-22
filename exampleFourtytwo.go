package main

import (
	"fmt"
	"os"
)

func FourtyTwo() {

	f:= createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}


func createFile(p string) *os.File {
	fmt.Println("Creating File")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f * os.File) {
	fmt.Println("writing file")
	fmt.Fprintln(f, "Data")
} 
func closeFile( f *os.File) {
	fmt.Println("Closing operation")
	f.Close()
}
