package main

import (
	"fmt"
	"time"
	"reflect"
)

func main()  {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month time.Month = now.Month()

	fmt.Println(reflect.TypeOf(month))
	fmt.Println("The", month, "month.")
	fmt.Println("in the", year, "year.")
}