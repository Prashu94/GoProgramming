package main

import "fmt"

func multipleUnnamed() {
	// calling a function an receiving multiple values from return statements
	fetchedName, fetchedDetails := setCustomerDetails("John-8097590094")
	fmt.Println(fetchedName, fetchedDetails)
}

func mulipleNamed() {
	fetchedName, fetchedDetails := setCustomerDetails1("John-8097590095")
	fmt.Println(fetchedName, " : ", fetchedDetails)
}

func setCustomerDetails1(details string) (customerDetails string, address string) {
	customerDetails = details
	address = "Street 11, NewYork"
	return
}

/*
	func setCustomerDetails(details string) (customerDetails string, string){
		if executed will throw on ERROR: mixed named and unnamed parameters
	}
*/
func setCustomerDetails(details string) (string, string) {
	customerDetails := details
	address := "Street 11, NewYork"
	return customerDetails, address // return multiple values from a single return statement
}

func main() {
	fmt.Println("Multiple Unnamed return values: ")
	multipleUnnamed()
	fmt.Println("Multiple Named return values: ")
	mulipleNamed()
}
