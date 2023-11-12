package main

import (
	"fmt"
	"log"
	"math/rand"
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
	d.shuffle()
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
	deckByteSliced:= d.toByteSlice()
	err := os.WriteFile(filename, deckByteSliced, 0666)
	if err != nil {
		return err
	}
	return nil
}

// newDeckFromFile reads a deck from a file specified by the filename.
// It attempts to open and read the file, returning an error if any issues arise.
// If successful, it splits the file content by new lines to create a deck of cards.
// Returns the new deck and nil if no error occurs.
func newDeckFromFile(filename string) (deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file '%s': %v ", filename, err)
		return nil, err
	}
	d := strings.Split(string(data), "\n")
	return deck(d), nil
}

// shuffle randomizes the order of cards in the deck.
// It iterates over the deck, swapping each card with another at a random position.
// This is achieved using the math/rand Intn function to generate random indices.
// The function modifies the deck in place and does not return any value.
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d [newPosition] = d[newPosition], d[i]
	}
}


