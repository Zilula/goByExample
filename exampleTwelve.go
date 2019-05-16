package main 

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}

func Twelve() {
	result1 := plus(1, 5)
	result2  := plusPlus(4, 6, 7)
	fmt.Println(result1, result2)
}