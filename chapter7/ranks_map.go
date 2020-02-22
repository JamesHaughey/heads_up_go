package main

import (
	"fmt"
)

func main() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	elements := map[string]string{
		"H":  "Hyrdrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
}
