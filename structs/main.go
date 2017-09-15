package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	sally := person{
		firstName: "Sally",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "ssmith@gmail.com",
			zipCode: 16009,
		},
	}

	sally.updateName("Samantha")
	sally.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
