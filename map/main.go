package main

import "fmt"

func main() {
	// All keys are strings; all values are strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	fmt.Println(colors)
}
