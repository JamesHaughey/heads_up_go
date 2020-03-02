package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}

	err := coordinates.SetLattitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

	location := geo.Landmark{}
	err = location.SetName("The googleplex")
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLattitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
