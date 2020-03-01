package main

import (
	"fmt"
)

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m * 1000)
}

func (g *Gallons) Double() {
	*g *= 2
}

func main() {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(10.0)
	busFuel = Liters(24.0)

	fmt.Println(carFuel, busFuel.ToGallons())

	carFuel.Double()
	fmt.Println("Doubled Fuel:", carFuel)

	truckFuel := Gallons(20)
	truckFuel.Double()
	fmt.Println("Truck Fuel:", truckFuel)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", soda, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", soda, milk.ToMilliliters())
}
