package main

import "fmt"

func zeroVal( ival int) {
	ival = 0
}

func zeroPtr(iptr * int) {
	 *iptr = 0
}

func SevenTeen() {
	i := 1
	fmt.Println("initial: ", i)
	zeroVal(i)
	fmt.Println("zeroVal: ", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr: ", i)
	fmt.Println("Pointer:", &i)
}

