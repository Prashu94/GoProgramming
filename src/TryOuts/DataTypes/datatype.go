package main

import (
	"fmt"
	"strconv"
)

func main() {
	id := 123
	fmt.Printf("'id' is of type '%T'\n", id)

	var uid uint32 = 213
	fmt.Printf("'uid' is of type '%T'\n", uid)

	// uncomment the below code and observe the output -- Compilation error
	//var uid uint32 = -21

	name := "12"
	fmt.Printf("'name' is of type '%T'\n", name)

	salary := 15000.68
	fmt.Printf("'salary' is of type '%T'\n", salary)

	var door string = "121"
	doorNo, _ := strconv.Atoi(door) //Atoi returns two values, we ignore the 2nd value by using '_'
	fmt.Printf("'door' is of type '%T' & 'doorNo' is of Type '%T'\n", door, doorNo)
	idFloat := float32(id) //type casting data type of id from 'int' to 'float32'
	fmt.Printf("'idFloat' is of type '%T'\n", idFloat)
}
