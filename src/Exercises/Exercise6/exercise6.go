package main

import "fmt"

func average(nums [6]int) {
	sum := 0
	for _, element := range nums {
		sum += element
	}
	fmt.Printf("Average of the marks: %d\n", sum/len(nums))
	average := sum / len(nums)
	count_equal := 0
	count_less := 0
	count_greater := 0
	for index := 0; index < len(nums); index++ {
		if average == nums[index] {
			count_equal++
		} else if average > nums[index] {
			count_less++
		} else if average < nums[index] {
			count_greater++
		}
	}
	if count_equal != 0 {
		fmt.Printf("Count of Marks equal to average is: %d\n", count_equal)
	}
	if count_less != 0 {
		fmt.Printf("Count of Marks less than average is: %d\n", count_less)
	}
	if count_greater != 0 {
		fmt.Printf("Count of Marks greater than average is: %d\n", count_greater)
	}
}
func main() {
	nums := [6]int{73, 78, 70, 83, 95, 69}
	average(nums)
}
