package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Alex"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
