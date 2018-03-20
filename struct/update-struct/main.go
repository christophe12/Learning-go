package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	// updating structs
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Printf("%+v\n", alex)
}
