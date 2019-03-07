package main

func main() {
	// var card string = "Aces of Spades"
	cards := deck{newCard(), newCard()}
	cards = append(cards, "The joke")

	cards.print()
}

func newCard() string {
	return "Five of Spades"
}
