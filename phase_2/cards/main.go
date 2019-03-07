package main

import "fmt"

func main() {
	// var card string = "Aces of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Spades"
}
