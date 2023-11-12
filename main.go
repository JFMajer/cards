package main

func main() {
	cards := newDeck()
	hand1, remainingCards := deal(cards, 5)
	hand1.print()
	hand1.saveToFile("hand2.txt")
	remainingCards.print()
}
