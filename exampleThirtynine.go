package main

import (
	"fmt"
	"sort"
)
func ThirtyNine() {
	strs := []string{"a", "f", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{6, 3, 6, 8, 1}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	s:= sort.IntsAreSorted(ints)
	fmt.Println("Sorted:  ", s)



}