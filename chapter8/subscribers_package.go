package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func showStatus(sub *magazine.Subscriber) {
	fmt.Printf("Name: %s, Active: %t, Rate: %0.2f\n", sub.Name, sub.Active, sub.Rate)
}

func applyDiscount(sub *magazine.Subscriber) {
	sub.Rate = 4.99
}

func minimumSubscriber(name string) *magazine.Subscriber {
	var sub magazine.Subscriber
	sub.Name = name
	sub.Rate = 100
	sub.Active = true
	return &sub
}

func main() {
	subscriber1 := &magazine.Subscriber{Name: "Aman Singh", Rate: 10, Active: true}
	subscriber2 := minimumSubscriber("Beth Ryan")
	subscriber2.Active = false

	subscriber3 := minimumSubscriber("George Harrison")
	applyDiscount(subscriber3)

	showStatus(subscriber1)
	showStatus(subscriber2)
	showStatus(subscriber3)
}
