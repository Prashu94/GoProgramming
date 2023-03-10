package main

import "fmt"

func shorthand() {
	// declaring and initializing variable greeting
	greeting := 20
	// declaring and initialization pointer, ptrGreeting
	ptrGreeting := &greeting

	fmt.Println("value of greeting", greeting)
	fmt.Println("address of greeting", &greeting)
	fmt.Println("value of greeting", ptrGreeting)

}

func builtIn() {
	// declaring and initializing variable greeting
	var greeting string = "Hello World"
	// declaring string pointer, ptrGreeting
	ptrGreeting := new(string)
	// initialization of pointer
	ptrGreeting = &greeting
	fmt.Println("value of greeting", greeting)
	fmt.Println("address of greeting", &greeting)
	fmt.Println("value of greeting", ptrGreeting)
}

func derefernce() {
	var greeting string = "Hello World"
	var ptrGreeting *string = &greeting
	fmt.Println("value of greeting ->", greeting)
	fmt.Println("address of greeting ->", &greeting)
	fmt.Println("value of ptrGreeting ->", ptrGreeting)
	fmt.Println("accessing value of greeting using pointer ptrGreeting ->", *ptrGreeting)
}

func main() {
	// declaring and initializing variable greeting
	var greeting string = "Hello World"

	// declaring string pointer, ptrGreeting
	var ptrGreeting *string
	// initalization of pointer
	ptrGreeting = &greeting

	fmt.Println("value of greeting", greeting)
	fmt.Println("address of greeting", &greeting)
	fmt.Println("value of ptrGreeting", ptrGreeting)

	shorthand()
	builtIn()
	derefernce()

	var ptr *string
	fmt.Println("value of ptr ->", ptr)
}
