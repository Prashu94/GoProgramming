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

func employeeName1(emp string) {
	emp = "Jack"
}

func employeeName2(emp *string) {
	*emp = "Jack"
}

func passByValue(a, b int) {
	temp := 0
	temp = a

	a = b
	b = temp
}

func passByReference(a, b *int) int {
	var o int

	o = *a
	*a = *b
	*b = o

	return o
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

// passing and returning pointers from functions
func change(value *int) {
	*value = 55
}

// function that returns the pointer as output
func method() *int {
	value := 10
	return &value
}

func main() {
	fmt.Println("Multiple Unnamed return values: ")
	multipleUnnamed()
	fmt.Println("Multiple Named return values: ")
	mulipleNamed()

	emp := "Daniel Abrhaim "
	fmt.Println("value emp before function call -> ", emp)
	employeeName1(emp)
	fmt.Println("value of emp after function call -> ", emp)
	employeeName2(&emp)
	fmt.Println("value of emp after function call to employee2 -> ", emp)

	var x int = 10
	var y int = 20
	fmt.Printf("x = %d and y = %d", x, y)
	passByValue(x, y)
	fmt.Printf("\nx = %d and y = %d", x, y)
	passByReference(&x, &y)
	fmt.Printf("\nx = %d and y = %d", x, y)

	a := 60
	fmt.Println("\nValue of a before change method call is", a)
	b := &a
	change(b)
	fmt.Println("Value of a after change method call is", a)

	num := method()
	fmt.Println("value of num is :", *num)
	fmt.Println("Value is : ", *method())
}
