package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIfWeCanPrintADeck(t *testing.T) {
	var d = deck{"Ace of Spades", "Three of Clubs"}
	d.print()
}

func TestToString(t *testing.T) {
	var d = newDeck()
	fmt.Println(d.toString())
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Invalid Deck, Expected 52 but got %v", len(d))
	}
}

func TestFirstAndLastCard(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Invalid Deck. First card should be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Invalid Deck. Last card should be King of Clubs, but got %v", d[len(d)-1])
	}
}

func cleanup(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("Error Cleaning up file %v:  %v", fileName, err)
	}
}

func TestDeckWriteRead(t *testing.T) {
	const filename = "_decktesting"
	cleanup(filename)
	d := newDeck()
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != len(d) {
		t.Errorf("Invalid Deck, Expected a length of %v, got %v", len(d), len(loadedDeck))
	}
	cleanup(filename)
}
