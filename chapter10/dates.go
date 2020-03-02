package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 0 {
		return errors.New("Invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}

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

	fmt.Println(date)
}
