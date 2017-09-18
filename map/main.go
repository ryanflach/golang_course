package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Map vs Struct
// Map:
// - All keys must be the same type
// - All values must be the same type
// - Keys are indexed and can be iterated over
// - Generally used to represent a collection of related properties
// - Don't need to know all the keys at compile time
// - Reference Type (passed as a reference, rather than a copy of the value)
// Struct:
// - Values can be of different type
// - Keys don't support indexing
// - You need to know all the different fields at compile time
// - Generally used to represent a 'thing' with a lot of different properties
// - Value Type
