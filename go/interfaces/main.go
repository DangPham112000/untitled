package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	enbo := englishBot{}
	spbp := spanishBot{}

	printGreeting((enbo))
	printGreeting((spbp))
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very different implement
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
