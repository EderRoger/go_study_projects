package main

import "fmt"

func main() {
	// var card string = "Aces of Spades"
	cards := []string{newCard(), newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Spades"
}
