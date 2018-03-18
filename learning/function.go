package main

import (
	"fmt"
)

func main() {
	card := newCard()

	fmt.Println(card)

}

// how to create a function
func newCard() string {
	return "Ace of shades"
}
