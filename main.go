package main

import "log"

func main() {
	cards := newDeck()
	hand1, remainingCards := deal(cards, 5)
	// hand1.print()
	err := hand1.saveToFile("hand1.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = remainingCards.saveToFile("remaining.txt")
	// deck, _ := newDeckFromFile("hand1.txt")
	// deck.print()
}
