package main

import "fmt"

func main() {
	// var card string = "Aces of Spades"
	cards := newDeck()
	hands, remainingDeck := deal(cards, 5)

	hands.print()
	fmt.Println("--------------------------------------------------------")
	remainingDeck.print()
}
