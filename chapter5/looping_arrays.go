package main

import (
	"fmt"
)

func main() {
	notes := [7]string{"Do", "Re", "Mi", "Fa", "So", "La", "Ti"}

	for i := 0; i < len(notes); i++ {
		fmt.Println("Note:", notes[i])
	}

	for idx, note := range notes {
		fmt.Println("Iteration:", idx, "Note:", note)
	}
}
