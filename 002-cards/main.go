package main

import "fmt"

func main() {
	deckSize := 52
	card := newCard()
	fmt.Println(deckSize, card)
}

func newCard() string {
	return "Five of Diamonds"
}
