package main

import "fmt"

// Approach 1
const customerId string = "Jack123"
const customerPassword = "jack@123"

// Approach 2
const (
	customerNumber = "8097590093"
	address        = 10
	AVAILABILITY   = true
)

func main() {
	fmt.Println(customerId)
	fmt.Println(customerNumber)
}
