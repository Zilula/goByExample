package main 

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func Fourteen() {
	sum(1, 4, 5)
	sum(1, 2, 3, 4, 5, 6)

	
	nums := []int{1, 2, 3, 4, 6}

	sum(nums...)
}