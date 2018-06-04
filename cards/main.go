package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// remainingCards.saveToFile("cards_game")

	//cards := newDeckFromFile("cards_game")
	//cards.print()

	// cards := newDeck()
	// cards.shufle()
	// cards.print()

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("is even", number)
		} else {
			fmt.Println("is odd", number)
		}
	}
}
