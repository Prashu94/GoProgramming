package main

import "fmt"

func swap(x *int, y *int) {
	var t int

	t = *x
	*x = *y
	*y = t
}

var x int
var y int

func main() {
	fmt.Scan(&x)
	fmt.Scan(&y)

	swap(&x, &y)

	fmt.Println(x)
	fmt.Println(y)
}
