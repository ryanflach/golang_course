package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	sally := person{
		firstName: "Sally",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "ssmith@gmail.com",
			zipCode: 16009,
		},
	}

	fmt.Printf("%+v\n", sally)
}
