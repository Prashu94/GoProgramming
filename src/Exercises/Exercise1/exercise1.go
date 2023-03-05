package main

import (
	"fmt"
)

// Enum Constants
type Discount int

const (
	_ Discount = iota
	new
	regular
	special
)
const percentage float64 = 0.05

var (
	name = "Joy"
	id   = 123
	cost = 2500
)

func main() {
	discountPrice := percentage * float64(regular)
	discountCost := float64(cost) - (float64(cost) * discountPrice)
	fmt.Printf("%v%v: %v", name, id, discountCost)
}
