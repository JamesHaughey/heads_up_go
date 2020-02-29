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

	primes := make(map[int]bool)
	primes[4] = false
	primes[7] = true
	fmt.Println(primes[7])
	fmt.Println(primes[4])
}
