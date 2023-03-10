package main

import "fmt"

func main() {
	type employee struct {
		employeeId   int
		employeeName string
	}

	emp := &employee{
		employeeId:   1,
		employeeName: "Harry",
	}

	fmt.Println("Employee Id: ", emp.employeeId)
	fmt.Println("Employee Name: ", emp.employeeName)
}
