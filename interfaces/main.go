package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// For demonstration purposes, assume this method has custom logic in it
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// For demonstration purposes, assume this method has custom logic in it
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
