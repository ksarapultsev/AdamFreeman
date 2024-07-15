package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Value:", rand.Intn(100))
	fmt.Println("Value Round:", math.Round(34.6))
	fmt.Println("Value Ceil:", math.Ceil(27.1))
	val1 := "true"
	val2 := "false"
	val3 := "truew"
	val4 := "500"
	val5 := "100"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)

	bool4, b4err := strconv.ParseInt(val4, 0, 8)
	bool5, b5err := strconv.ParseInt(val5, 0, 8)

	if b4err == nil {
		fmt.Println("Parsed value:", bool4)
	} else {
		fmt.Println("Cannot parse", val4, b4err)
	}

	if b5err == nil {
		smallInt := int8(bool5)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val5, b5err)
	}
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)

	val6 := "48.95"
	float1, float1err := strconv.ParseFloat(val6, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}
}
