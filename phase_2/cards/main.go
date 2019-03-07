package main

import "fmt"

func main() {
	// var card string = "Aces of Spades"
	cards := []string{newCard(), newCard()}
	cards = append(cards, "The joke")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Spades"
}
