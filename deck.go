package main

import (
	"fmt"
	"os"
	"strings"
)

// Slice of string representing card suits
var cardSuits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}

// Slice of string representing card values
var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

type deck []string

// newDeck creates and returns a new deck of playing cards.
// It combines each card value with every suit to create a standard deck.
func newDeck() deck {
	d := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%s of %s", value, suit)
			d = append(d, card)
		}
	}
	return d
}


// print displays each card in the deck to the standard output.
// It iterates through the deck and prints every card's description.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal splits the deck into two parts: a hand and the remaining deck.
// It takes a deck and a hand size, returning two decks (the hand and the rest).
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toByteSlice converts the deck into a byte slice.
// This is primarily used for file operations, like saving the deck to a file.
func (d deck) toByteSlice() []byte {
	joinedString := strings.Join(d, "\n")
	byteSlice := []byte(joinedString)
	return byteSlice
}

// saveToFile saves the deck to a file specified by the filename parameter.
// It converts the deck into a byte slice and writes it to the file.
func (d deck) saveToFile(filename string) error {
	deckStringified := d.toByteSlice()
	err := os.WriteFile(filename, deckStringified, 0666)
	if err != nil {
		return err
	}
	return nil
}

func readFromFile(filename string) (deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	d := strings.Split(string(data), "\n")
	return deck(d), nil
}


