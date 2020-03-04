package main

import (
	"fmt"

	"github.com/headfirstgo/gadget"
)

// This playlist just works with TapePlayer structs
func playList(device gadget.TapeInterface, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player gadget.TapeInterface) {
	player.Play("Test Track")
	player.Stop()

	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Type Assertion:", ok)
	}
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player2 := gadget.TapeRecorder{}
	mixtape2 := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player2, mixtape2)

	var tapeInterface1 gadget.TapeInterface = gadget.TapePlayer{}
	var tapeInterface2 gadget.TapeInterface = gadget.TapeRecorder{}

	TryOut(tapeInterface1)
	TryOut(tapeInterface2)
}
