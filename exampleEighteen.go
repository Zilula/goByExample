package main

import "fmt"

type person struct {
	name string
	age int
}

func Eighteen() {
	fmt.Println(person{"Lance", 22})
	fmt.Println(person{name: "Zach", age: 26})
	fmt.Println(person{name: "Jack"})
	fmt.Println(&person{name: "lance2", age: 44})

	s := person{name: "sean", age: 43}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)

	sp.age = 5432
	fmt.Println(sp.age)


}