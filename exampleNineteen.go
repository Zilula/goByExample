package main

import "fmt"

type rect struct {
	width, height int
}

func (banana rect) area() int {
	return banana.width * banana.height
}
func (apple rect) perim() int {
	return 2*apple.height + 2*apple.width
}

func nineteen() {
	banana := rect{width: 12, height: 4}
	fmt.Println("area", banana.area())
	fmt.Println("perim", banana.perim())

	copyOfBanana := &banana

	fmt.Println("area: ", copyOfBanana.area())
	fmt.Println("perim: ", copyOfBanana.perim())
	copyOfBanana.width = 14
	fmt.Println(banana)

}
