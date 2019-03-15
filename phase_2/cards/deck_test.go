package main

import "testing"

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
