package main

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 54, but got %v", len(d))
	}
}

func TestDeal(t *testing.T) {
	cards := newDeck()
	hand1, remainingCards := deal(cards, 5)

	if len(hand1) != 5 {
		t.Errorf("Expected hand size of 5, but got %v", len(hand1))
	}

	if len(remainingCards) != 47 {
		t.Errorf("Expected remaining deck size of 47, but got %v", len(remainingCards))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	dOriginal := make(deck, len(d))
	copy(dOriginal, d)
	d.shuffle()

	if reflect.DeepEqual(d, dOriginal) {
		t.Errorf("Decks are the same after shuffle operation")
	}
}
