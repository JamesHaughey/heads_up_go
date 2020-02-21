package main

import (
	"fmt"
)

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	// Change the underlying array and the values of the slice will change
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice2 := array1[0:3]

	fmt.Println("Slice value before change:", slice2)
	array1[1] = "X"
	fmt.Println("Slice value after change:", slice2)
	fmt.Println("Array value after change:", array1)

}
