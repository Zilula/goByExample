package main

import "fmt"

func inSeq() func() int {
	i := 0 
	return func()  int {
		i ++
		return i
	}
}
// "Fiftheen is used in main.go"
func Fifteen() {
	nextInt := inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := inSeq()
	fmt.Println(newInts())
}