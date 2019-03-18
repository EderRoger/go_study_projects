package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(d))
	}

	firstCard := d[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Exepected first card of Ace of Spades, but got %v", firstCard)
	}

	lastCard := d[len(d)-1]
	if lastCard != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDesk := newDeckFromFile("_decktesting")

	if len(loadedDesk) != 16 {
		t.Errorf("Expected 16 cards, but got, %v", len(loadedDesk))
	}

	os.Remove("_decktesting")
}
