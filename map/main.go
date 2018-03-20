package main

import "fmt"

func main() {
	// first way of declaring a map
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	// second way of declaring a map
	//var cars map[string]string

	// third way of declaring a map
	//bus := make(map[int]string)

	// adding an element to an existing map
	colors["black"] = "#000000"

	// removing an element to an existing map
	delete(colors, "red")

	printMap(colors)
}

// iterating over a map
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex for color ", key, " is ", value)
	}
}
