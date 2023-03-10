package main

import "fmt"

func tryout1() {
	// creating a pointer
	empName := "Rachel Green!"
	var ptrEmp = &empName
	fmt.Printf("Type of ptrEmp is %T\n", ptrEmp)
	fmt.Println("Address of empName is", &empName)

	// Zero value of a pointer is nil
	age := 25
	var ptrAge *int
	if ptrAge == nil {
		fmt.Println("\n age is", ptrAge)
		// initialize the pointer
		ptrAge = &age
		fmt.Println("ptrAge after initialization is", ptrAge)
	}

	// dereferencing a pointer
	salary := 25500
	ptrSalary := &salary
	fmt.Println("\nAddress of salary is", ptrSalary)
	fmt.Println("Value of salary is", *ptrSalary)

	// changing the value pointed using reference
	*ptrSalary++
	fmt.Println("\nNew value of salary is", salary)
}

func tryout2() {
	type employee struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}

	emp := &employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name: ", &emp.firstName)
	fmt.Println("Age: ", &emp.age)
	fmt.Println("First Name: ", (*emp).firstName)
	fmt.Println("Age: ", (*emp).age)
}

func main() {
	tryout1()
	tryout2()
}
