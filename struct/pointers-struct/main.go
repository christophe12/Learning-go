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

	alex.updateProperty("Alexia")

	alex.print()

}

func (p *person) updateProperty(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

/**
* Take aways
*
* For these types no need to use pointers when updating
*	SLICES
*	MAPS
*	CHANNELS
*	POINTERS
*	FUNCTIONS
*
* For these types, use pointers when updating
*	INT
*	FLOAT
*	STRING
*	BOOL
*	STRUCT
*
 */
