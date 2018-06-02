package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// remainingCards.saveToFile("cards_game")

	//cards := newDeckFromFile("cards_game")
	//cards.print()

	cards := newDeck()
	cards.shufle()
	cards.print()
}
