package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // another way to declare a struct by omitting the struct var name
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 21000,
		},
	}

	alex.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
