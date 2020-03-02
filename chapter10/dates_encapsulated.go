package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	var date calendar.Date

	err := date.SetYear(1910)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Year:", date.Year())
	fmt.Println("Month:", date.Month())
	fmt.Println("Day:", date.Day())

	var birthday calendar.Event

	err = birthday.SetYear(1984)
	if err != nil {
		log.Fatal(err)
	}

	err = birthday.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}

	err = birthday.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}

	err = birthday.SetTitle("James' Birthday that is really really really really really really really long")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(birthday.Title())
}
