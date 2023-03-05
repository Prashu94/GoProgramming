package main

import "fmt"

func main() {

	customerBill := 340.0
	fixedCharges := 20.0
	oneDayCharge := 10.0
	discount := 20.0

	totalCustomerBill := customerBill + fixedCharges // Addition
	fmt.Printf("Total Bill for you is = Rs. %v", totalCustomerBill)

	FixedCharges := totalCustomerBill - customerBill // Subtraction
	fmt.Printf("\nFixed Charges applied to your plan are Rs. %v", FixedCharges)

	oneMonthBill := oneDayCharge * 30 // Multiplication
	fmt.Printf("\nYour One Month Bill is = Rs. %v", oneMonthBill)

	billAfterDiscount := oneMonthBill - (oneMonthBill*discount)/100 // Division
	fmt.Printf("\nBill After Discount = Rs. %v", billAfterDiscount)

	moduleExample := 40 % 3 // Modulus
	fmt.Printf("\nResult of 40%%3 = %v", moduleExample)

	// CASE 1: Result_Type1 = Type1_variable <operator> Type1_variable2:
	var count int = 25
	var flag int = 32
	result01 := count * flag
	fmt.Printf("%T\n", result01)

	// CASE 2: Result_Type1 = Type1_variable <operator> Untyped_variable2
	var count2 float64 = 12.1
	flag2 := 31.3
	result02 := flag2 - count2
	fmt.Printf("%T\n", result02)

	// CASE 3: Result_Type = Untyped_variable1 <operator> Untyped_variable2
	count3 := 31
	const flag3 = 'A'
	result03 := count3 * flag3
	fmt.Printf("%T\n", result03)
}
