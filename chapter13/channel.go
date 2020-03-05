package main

import (
	"fmt"
	"time"
)

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
	fmt.Println("ABC Load finished")
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
	fmt.Println("DEF Load finished")
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)

	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
	time.Sleep(time.Second * 5)
}
