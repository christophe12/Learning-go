package main

import "fmt"

func main() {

	numbersSlice := []int{}

	// adding numbers to slice
	counter := 0

	for counter <= 10 {
		numbersSlice = append(numbersSlice, counter)
		counter = counter + 1
	}

	for _, number := range numbersSlice {
		if (number % 2) == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}

}
