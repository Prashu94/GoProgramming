package main

import "fmt"

var customerLoginId int = 2300 // Standard Way
var validity int               // Initializing later
var years = 23                 // Type Inference

// Variable Declaration Block
var (
	customerName    = "Prashant"
	customerAge     = 50
	customerContact = "8094590093"
)

// variable declared at package level
var customerLogin string = "C1001"

var validityUser int = 100

// Global Level ---> Accessible by any package and function directly
var CustomerName string = "John"

// Package Level ---> Accessible by all functions in this package directly
var customerContact1 string = "80959003"

func main() {
	fmt.Println(customerLoginId, validity, years)

	fmt.Println(customerName, customerAge, customerContact)
	// Variable declaration block at the function level
	var (
		planValidity int = 25
		counter      int
		totalCounts  = 200
	)

	fmt.Println(planValidity)
	counter = 21
	i, j, k := 1, 2, 3
	fmt.Println(counter, totalCounts)
	fmt.Println(i, j, k)

	fmt.Println(customerLogin)
	// variable redeclared at the block level
	var customerLogin string = "C1002"
	// reinitializing will shadow away previous value
	validityUser = 100
	fmt.Println(customerLogin)
	// COMPILATION ERROR: customerLogin redeclared
	//var customerLogin = 200
	parse()
	fmt.Println(customerLogin)

	// Block Level ---> Accessible inside this function block only
	var customerLastName string = "Tait"

	customerFullName := CustomerName + customerLastName
	fmt.Printf("%v \nContact No. %v ", customerFullName, customerContact1)
}

func parse() {
	customerLogin = "C202"
	fmt.Println(customerLogin)
}
