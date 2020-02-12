package main

import (
	"fmt"
	"log"
)

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i:= 0; i < times; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10)
}

func paintNeededReturn(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func paintNeededReturnError(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	sayHi()

	repeatLine("This is repetitive!", 3)

	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)

	fmt.Println()
	var amount, total float64
	amount = paintNeededReturn(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = paintNeededReturn(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)

	fmt.Println()
	
	amount, err := paintNeededReturnError(4.2, -3.0)
	if err != nil {
		log.Fatal(err)	
	} else {
		fmt.Printf("%0.2f liters neeeded\n", amount)
	}
}