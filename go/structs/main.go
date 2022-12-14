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

	// danhPointer := &danh
	// danhPointer.updateName("Danheeeee")
	// (&danh).updateName("Danheeeee&")
	danh.updateName("Danheeeeeh")
	danh.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
