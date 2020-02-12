package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func main() {
	var myInt int
	var myIntPointer *int

	myInt = 42
	myIntPointer = &myInt

	fmt.Println(*myIntPointer)
	fmt.Println()

	amount:= 6
	fmt.Println("Amount is:", amount)

	double(&amount)
	fmt.Println("Doubled Amount is:", amount)
}