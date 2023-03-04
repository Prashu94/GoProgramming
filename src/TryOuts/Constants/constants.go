package main

import (
	"fmt"
)

// Package Level Constants
const pi float32 = 3.14

func main() {
	//Constants declaration inside main function
	//	const pi float = 3.14159
	const courseName string = "JavaScript"
	fmt.Println("Constant Declaration inside main function: ")
	fmt.Println(courseName)

	fmt.Println("Constant declaration at Package level: ")
	fmt.Println(pi)
}
