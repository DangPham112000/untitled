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
	danh := person{
		firstName: "Danh",
		lastName:  "Phan",
		contact: contactInfo{
			email:   "pktdanh@gmail.com",
			zipCode: 70000,
		},
	}

	danh.updateName("Danh lol")
	danh.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
