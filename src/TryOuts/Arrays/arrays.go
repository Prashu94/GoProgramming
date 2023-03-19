package main

import "fmt"

func tryout1() {
	// declaring an array named manufacturers that stores the details (id and checked) of manufacturers
	manufacturers := [5]string{"Samsung", "Motorola", "Apple", "Sony", "Toshiba"}

	// USE of for loop
	fmt.Println("1. Available Manufacturers using Simple For Loop: ")
	// iterates over the manufacturers array to display the id of all the manufactuers
	for index := 0; index < len(manufacturers); index++ {
		fmt.Println("Index --> ", index, "Manufacturer at index --> ", manufacturers[index])
	}

	// USE OF for .. of loop
	fmt.Println("2. Avaialble Manufacturers using: range keyword of for loop")
	// iterates over the manufactures array to display all the manufacturer details
	for _, value := range manufacturers {
		fmt.Println(value)
	}

	// Print the third element of the manufacturers array using index
	fmt.Println("3. Third element of the array: ", manufacturers[2])

	// Replace the last element of the array with "Google"
	manufacturers[len(manufacturers)-1] = "Google"
	fmt.Println("4. Replacing 'Toshiba' with 'Google'")
	fmt.Println(manufacturers)
}

func tryout2() {
	const noOfCustomers int = 3
	const noOfItems int = 4

	// initializing a multi-dimensional array
	costOfPurchase := [noOfCustomers][noOfItems]int{{540, 280, 270, 100}, {250, 600, 256, 178}, {115, 280, 123, 540}}

	fmt.Println(costOfPurchase)

	finalAmount := 0
	discount := 15

	// traversing the array using for loop
	for i := 0; i < noOfCustomers; i++ {
		sum := 0
		for j := 0; j < noOfItems; j++ {
			sum += costOfPurchase[i][j]
		}
		fmt.Print("Cost of purchase :", sum, "...")
		if sum >= 1000 {
			finalAmount = sum - (sum * discount / 100)
			fmt.Println("You are eligible for 15% discount!! Final Amount: ", finalAmount)
		} else {
			finalAmount = sum + 5
			fmt.Println("You have to pay extra tax amount of 5! Final Amount :", finalAmount)
		}
	}
}

func main() {
	tryout1()
	tryout2()
}
