package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}

	for _, line := range lines {
		counts[line]++
	}
	for key, value := range counts {
		fmt.Printf("Votes for %s: %d\n", key, value)
	}
}
