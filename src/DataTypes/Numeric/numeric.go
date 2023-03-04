package main

import "fmt"

func main() {

	// Integer
	var planPrice uint8 = 225 // Using 8-bit unsigned int
	fmt.Println(planPrice)

	var validityDays int16 = 32767 // Using 16-bit signed int
	fmt.Println(validityDays)

	// Floating
	var averageSales float32 = 10
	quarterlyAverage := 20.1 // type of quaterlyAverage is float64
	fmt.Printf("averageSales type %T, quarterlyAverage type %T", averageSales, quarterlyAverage)

	// Complex
	var num1 complex128 = 2 + 3i        // creating a complex type test variable
	var num2 complex64 = complex(2, 10) // creating complex type variable using inbuilt function

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("%v, %T \n", real(num1), real(num1)) //using inbuilt function to render real part
	fmt.Printf("%v, %T \n", imag(num1), imag(num1)) //using inbuilt function to render imaginary part

	// Runes
	arr1 := 'A'
	arr := "♠"
	fmt.Printf("Character %c, Unicode %U", arr1, arr1)
	fmt.Printf("\nCharacter %c, Unicode %U", arr, arr)
}
