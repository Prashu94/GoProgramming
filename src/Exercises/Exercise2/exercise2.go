package main

import (
	"fmt"
	"math"
)

func armstrong(num1 int) {
	n1 := num1
	var sum int = 0
	for n1 > 0 {
		i := n1 % 10
		n1 = n1 / 10
		sum += int(math.Pow(float64(i), 3))
	}
	if sum == num1 {
		fmt.Printf("%d is an Amstrong number", num1)
	} else {
		fmt.Printf("%d is not an Amstrong number", num1)
	}
}

var num1 int

func main() {
	fmt.Scan(&num1)
	armstrong(num1)
}
