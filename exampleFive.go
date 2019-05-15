package main 

import "fmt"

func main() {
	i:=1
	for i <= 3 {
		fmt.Println("loop one")
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <=9; j++ {
		fmt.Println("loop two")
		fmt.Println(j)
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 2 {
			continue
		}
		fmt.Println("loop three")
		fmt.Println(n)
	}
}