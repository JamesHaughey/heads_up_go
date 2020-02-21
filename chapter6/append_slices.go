package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b"}
	fmt.Println("Before Change:")
	fmt.Println("Slice Contents:", slice, ", Slice Length:", len(slice))
	slice = append(slice, "c")
	fmt.Println("After Change One:")
	fmt.Println("Slice Contents:", slice, ", Slice Length:", len(slice))
	slice = append(slice, "d", "e")
	fmt.Println("After Change Two:")
	fmt.Println("Slice Contents:", slice, ", Slice Length:", len(slice))
	fmt.Println()

	fmt.Println("Add a value to an empty slice")
	var intSlice []int
	intSlice = append(intSlice, 27, 15)
	fmt.Printf("intSlice: %#v\n", intSlice)
}
