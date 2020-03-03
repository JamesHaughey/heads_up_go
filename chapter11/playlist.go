package main

import (
	"github.com/headfirstgo/gadget"
)

// This playlist just works with TapePlayer structs
func playList(device gadget.TapeInterface, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player2 := gadget.TapeRecorder{}
	mixtape2 := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player2, mixtape2)
}
