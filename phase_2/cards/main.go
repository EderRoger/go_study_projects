package main

func main() {
	fileName := "my_cards"

	cards := newDeck()
	cards.saveToFile(fileName)

	cardsFromFile := newDeckFromFile(fileName)
	cardsFromFile.print()

}
