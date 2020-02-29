package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showStatus(sub *subscriber) {
	fmt.Printf("Name: %s, Active: %t, Rate: %0.2f\n", sub.name, sub.active, sub.rate)
}

func applyDiscount(sub *subscriber) {
	sub.rate = 4.99
}

func minimumSubscriber(name string) *subscriber {
	var sub subscriber
	sub.name = name
	sub.rate = 100
	sub.active = true
	return &sub
}

func main() {
	subscriber1 := minimumSubscriber("Aman Singh")
	subscriber2 := minimumSubscriber("Beth Ryan")
	subscriber2.active = false

	subscriber3 := minimumSubscriber("George Harrison")
	applyDiscount(subscriber3)

	showStatus(subscriber1)
	showStatus(subscriber2)
	showStatus(subscriber3)
}
