package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	fmt.Printf("Variable i: Value : %v, Type: %T \n", i, i)

	var f32 float32 = 10.6556
	fmt.Printf("Variable f32: Value: %v, Type: %T \n", f32, f32)

	i16 := int16(i) // Converting int8 (by default) -> int16
	fmt.Printf("Variable i16: Value: %v, Type: %T \n", i16, i16)

	i32 := int32(i) // Converting int16 -> int32
	fmt.Printf("Variable i32: Value: %v, Type: %T \n", i32, i32)

	i64 := int64(i) // Converting int32 -> int64
	fmt.Printf("Variable i64: Value: %v, Type: %T \n", i64, i64)

	f64 := float64(f32) // Converting float32 -> float64
	fmt.Printf("Variable f64: Value: %v, Type: %T \n", f64, f64)

	var strVar string = strconv.FormatInt(i64, 10) // Converting int64 to string using FormatInt()
	// func(from "strconv" package)
	fmt.Printf("Variable strVar: Value: %v, Type: %T \n", strVar, strVar)

	strVar2 := "100"
	intVar, _ := strconv.Atoi(strVar2) // Converting string value stored int strVar to int using Atoi() func
	fmt.Printf("Variable intVar: Value: %v, Type: %T \n", intVar, intVar)

	floatVar, _ := strconv.ParseFloat(strVar, 8) //Converting string to float using ParseFloat func.
	fmt.Printf("Variable floatVar: Value: %v, Type: %T \n", floatVar, floatVar)
}
