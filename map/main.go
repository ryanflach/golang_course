package main

import "fmt"

func main() {
	// alt syntax
	// var colors map[int]string

	colors := make(map[int]string)

	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)
}
