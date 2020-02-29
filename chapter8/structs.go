package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 5.09
	myStruct.word = "Something"
	myStruct.toggle = true

	fmt.Println(myStruct.word)
}
