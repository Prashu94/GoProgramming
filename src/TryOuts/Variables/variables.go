package main

import "fmt"

// Package Level variable declaration
var custId int = 101 // Type 1

// Variable declaration using var block
var ( // Type 2
	penName  string = "xyz"
	penId    int    = 563
	penBrand string = "Pilot"
)

func main() {
	// variable declaration inside main function
	var totalMarks int = 52  // Type 1
	var averageMarks float64 // Type 2
	averageMarks = 45.0
	average := 48 // Type 3

	// Print statements to print the values
	fmt.Println("Variable declaration inside main method.")
	fmt.Println(totalMarks)
	fmt.Println(averageMarks)
	fmt.Println(average)

	// Variable declaration at package level
	fmt.Println("Variable declaration at package level: ")
	fmt.Println(custId)

	// Variable declaration using var block
	fmt.Println("Variable declaration using var block: ")
	fmt.Println(penName)
	fmt.Println(penId)
	fmt.Println(penBrand)

}
