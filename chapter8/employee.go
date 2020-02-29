package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000

	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	employee.Address = address
	fmt.Println(employee.City)
}
