package main

import (
	"fmt"
)

func main() {
	col_format := "%12s | %2d\n"
	
	fmt.Println("Formatting Numbers")
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------")
	fmt.Printf(col_format, "Stamps", 50)
	fmt.Printf(col_format, "Paper Clips", 5)
	fmt.Printf(col_format, "Tape", 99)
	fmt.Println("")

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}