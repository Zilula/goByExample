package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 = odd")
	}
	if 8%2 == 0 {
		fmt.Println("8 is divisiable by 4")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if  num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "greater than 10")
	}

}