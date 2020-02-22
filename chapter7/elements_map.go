package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"

	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
}
