package main

import (
	"fmt"
)

func main() {
	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)
	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	counter1 := 0
	for counter1 <= 3 {
		fmt.Println("Counter:", counter1)
		counter1++
		// if (counter>3){
		// 	break
		// }
	}

	names := []string{"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)
}
