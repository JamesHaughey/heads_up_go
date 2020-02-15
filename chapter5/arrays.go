package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7] string

	notes[0] = "Do"
	notes[1] = "Re"
	notes[2] = "Mi"

	fmt.Println("Note One:", notes[0])
	fmt.Println("Note Two:", notes[1])
	fmt.Println("Note Five:", notes[4])
	fmt.Println("Note Six:", notes[5])


	var primes [5] int

	primes[0] = 2
	primes[1] = 3

	fmt.Println(primes[1])

	var dates [3] time.Time

	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

	var fullNotes [7] string = [7] string {"Do", "Re", "Mi", "Fa", "So", "La", "Ti"}
	fmt.Println(fullNotes[3], fullNotes[6], fullNotes[0])

	fullPrimes := [5] int {2, 3, 5, 7, 11}
	fmt.Println(fullPrimes[0], fullPrimes[2], fullPrimes[4])

	text := [3] string {
		"This is a series of long strings,",
		"which would be awkward to place",
		"together on a single line",
	}

	fmt.Println(text[2])
}