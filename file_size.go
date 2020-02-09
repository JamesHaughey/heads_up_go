package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("test.exe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Size())
}