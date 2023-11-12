package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Slice of string representing card suits
var cardSuits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}

// Slice of string representing card values
var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

type deck []string

// function to create new deck
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


// function that prints entire deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toByteSlice() []byte {
	joinedString := strings.Join(d, "\n")
	byteSlice := []byte(joinedString)
	return byteSlice
}

func (d deck) saveToFile(filename string) {
	deckStringified := d.toByteSlice()
	err := os.WriteFile(filename, deckStringified, 0666)
	if err != nil {
		log.Fatal(err)
	}
}


