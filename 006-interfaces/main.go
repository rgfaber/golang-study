package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {
}
type spanishBot struct {
}
type frenchBot struct{}

func (f frenchBot) getGreeting() string {
	return "Bonjour!"
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	fb := frenchBot{}
	printGreeting(eb)
	printGreeting(sb)
	printGreeting(fb)
}
