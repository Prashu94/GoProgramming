package main

import "fmt"

func tryout1() {
	num1 := 10
	num2 := 20
	fmt.Println(add(num1, num2))
	fmt.Println(subtract(num1, num2))
}

func add(num3, num4 int) int {
	add := num3 + num4
	return add
}

func subtract(num3, num4 int) int {
	subtract := num3 - num4
	return subtract
}

func nums(nums ...int) {
	temp := 0
	for _, i := range nums {
		temp += i
	}
	fmt.Println(temp)
}

func tryout2() {
	nums(1, 2, 3, 4)
}

func tryout3() {
	num1 := 100
	num2 := 200
	fmt.Println("Before swap num1: ", num1, "\nBefore swap num2: ", num2)
	//swapRef(&num1, &num2)
	swapVal(num1, num2)
	fmt.Println("After swap num1: ", num1, "\nAfter swap num2: ", num2)
}

func swapRef(num3 *int, num4 *int) {

	var temp int
	temp = *num3
	*num3 = *num4
	*num4 = temp
}

func swapVal(num3, num4 int) {
	var temp int
	temp = num3
	num3 = num4
	num4 = temp
}

func tryout4() {
	// passing pointer to a function
	a1 := 58
	fmt.Println("\nValue of a1 before function call is", a1)
	ptrA1 := &a1
	change(ptrA1)
	fmt.Println("\nValue of a1 after function all is", a1)

	// slice as an argument to a function
	o := [3]int{89, 90, 91}
	modify(o[:])
	fmt.Println("modify using slice", o)
}

func change(age *int) {
	*age = 55
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	tryout1()
	tryout2()
	tryout3()
	tryout4()
}
