package main 

import "fmt"

func Eleven() {

	nums := []int{1, 4, 5, 7, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3 = ", i )
		}
	}
	keymap := map[string]string{"key1": "value 1", "keytwo": "valueTwo"}

	for k, v := range keymap {
		fmt.Printf("%s -> %s/n ", k, v)
	}

	for k :=  range keymap {
		fmt.Println("keys:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}