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

// Enum Constants
type Directory int // enum Directory

const (
	_        Directory = iota // throws away first iota value i.e = 0
	Contact1                  // Contact1 = 1
	Contact2                  // Contact2 = 2
	Contact3                  // Contact3 = 3
	Contact4                  // Contact4 = 4
)

func main() {
	fmt.Println(customerId)
	fmt.Println(customerNumber)

	// Declaring a variable NewContact with type Directory
	var NewContact Directory

	NewContact = 2
	//fmt.Println(NewContact == Contact3)
	fmt.Println(NewContact == Contact2)
}
