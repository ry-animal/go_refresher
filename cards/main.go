package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	deckFromFile := newDeckFromFile("my_cards")

	hand, _ :=deckFromFile.deal(5)

	fmt.Println("-------------------")
	hand.print()
	fmt.Println("-------------------")

}