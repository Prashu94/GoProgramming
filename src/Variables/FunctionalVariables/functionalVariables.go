package main

import "fmt"

func main() {

	// 1. Initializing at the time of declaration
	var customerId int = 2501
	fmt.Println(customerId)

	// 2. Initializing later (Useful while using loops)
	var validity int
	validity = 21

	fmt.Println(validity)

	// 3. Initializing and Declaration with type inference
	var LoginMessage = "Welcome User"
	fmt.Println(LoginMessage)

	// 4. Shorthand Initialization and Declaration
	price := 100.0
	fmt.Println(price)

	// 5. Initializing and Declaration of multiple variables
	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	// 6. Using blank identifier
	_ = 32

}
