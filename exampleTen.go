package main 

import "fmt"
func Ten() {
	m := make(map[string]int)

	m["keyOne"] = 4
	m["keyTwo"] = 55

	fmt.Println("map:", m)
	
	v1 := m["keyOne"]
	fmt.Println("v1:", v1)
	fmt.Println("length:", len(m))
	
	delete(m, "keyTwo")
	fmt.Println("map:", m)

	_, prs := m["keyTwo"]
	fmt.Println("prs:", prs)

	n := map[string]int{"one": 1, "two": 2}
	fmt.Println("map:", n)
}