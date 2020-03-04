package main

import (
	"fmt"
)

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main() {
	fmt.Println("The first jug contains", Liters(5.05498960))
	fmt.Println(Gallons(5.049864877))
	fmt.Println(Milliliters(50.245842566))
}
