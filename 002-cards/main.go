package main

import "fmt"

func main() {
	var cards = newDeck()
	hand, rest := deal(cards, 3)
	fmt.Println("Hand")
	hand.print()

	fmt.Println("Rest")
	rest.print()
	fmt.Println(hand.toString())
	err := hand.saveToFile("hand.txt")
	if err != nil {
	}
	myDeck := newDeck()
	err2 := myDeck.saveToFile("deck")
	if err2 != nil {
		return
	}
	readDeck := newDeck()
	readDeck.print()
	readDeck.shuffle()
	readDeck.print()
}
