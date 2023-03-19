package main

import "fmt"

func multipleReturns(num1, num2 int) (int, int, int, float32) {
	return num1 + num2, num1 - num2, num1 * num2, float32(num1 / num2)
}
func main() {
	fmt.Println(multipleReturns(10, 35))
}
