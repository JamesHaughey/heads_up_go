package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url, len(body))
}

func main() {
	go responseSize("https://www.trustedreviews.com")
	go responseSize("https://golang.org")
	go responseSize("https://golang.org/doc")
	time.Sleep(time.Second * 5)
	fmt.Println("End Main")
}