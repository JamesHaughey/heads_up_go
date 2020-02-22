package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}

	// This loop prints the grades in a random order
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.2f\n", name, grade)
	}

	// This loop prints grades in alphabetical order
	var names []string
	for name := range grades {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f\n", name, grades[name])
	}
}
