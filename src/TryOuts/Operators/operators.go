// Write a code to deduct 5.6% of tax on total interim salary
package main

import "fmt"

func main() {
	baseSalary := 15000 + 500.34 // baseSalary will take the data type of the latter, it untyped variable is included in the operation
	var internSal float32 = 1000
	totalSalary := baseSalary + float64(internSal)
	var bonus = 0
	var totalInternSal = (baseSalary * float64(bonus) / 100) + totalSalary
	fmt.Println("Total Internal Salary : ", totalInternSal)
	num := 10
	var value float32 = 30
	value = float32(num)/float32(value) + float32(num)*float32(num) - float32(value)
	fmt.Println(value)
}
