package main

import "log"

func main() {
	cards := newDeck()
	hand1, _ := deal(cards, 5)
	// hand1.print()
	err := hand1.saveToFile("hand2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remainingCards.print()
	deck, _ := readFromFile("hand1.txt")
	deck.print()
}
