package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showStatus(sub subscriber) {
	fmt.Printf("Name: %s, Active: %t, Rate: %0.2f\n", sub.name, sub.active, sub.rate)
}

func applyDiscount(sub *subscriber) {
	sub.rate = 4.99
}

func minimumSubscriber(name string) subscriber {
	var sub subscriber
	sub.name = name
	sub.rate = 100
	sub.active = true
	return sub
}

func main() {
	var subscriber1 subscriber
	var subscriber2 subscriber

	subscriber1.name = "Aman Singh"
	subscriber2.name = "Beth Ryan"
	subscriber1.active = true

	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Name:", subscriber2.name)

	showStatus(subscriber1)

	subscriber3 := minimumSubscriber("George Harrison")
	applyDiscount(&subscriber3)
	showStatus(subscriber3)
}
