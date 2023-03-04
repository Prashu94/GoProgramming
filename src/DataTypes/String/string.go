package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Strings
	loginPass := "Welcome User ! \n Login Successful" // creating using shorthand declaration
	var age string                                    // creating string using var keyword
	age = "27 years"
	fmt.Println("String 1: ", loginPass)
	//loginPass[2] = "No" // compilation error
	fmt.Println("String 2: ", age)

	loginPass1 := "Welcome User ! \n Login Successful" // creating interpreted string literals
	var age1 string
	age1 = `27 years \n` // \n won't be recognized by raw string literal

	fmt.Println("String 1: ", loginPass1)
	fmt.Println("String 2: ", age1)

	fmt.Println("String 1 Length is : ", len(loginPass)) //length calculated using len()

	fmt.Println("String 2 Length is : ", utf8.RuneCountInString(age)) //using RuneCountString
}
