package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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

func TestFileWrite(t *testing.T) {
	cleanUp()

	file1 := "hand_test.txt"
	file2 := "remaining_cards_test.txt"
	cards := newDeck()
	hand1, remainingCards := deal(cards, 5)
	err := hand1.saveToFile(file1)
	if err != nil {
		t.Fatal(err)
	}
	err = remainingCards.saveToFile(file2)
	if err != nil {
		t.Fatal(err)
	}

	if lines, err := countLines(file1); err != nil || lines != 5 {
		t.Errorf("File %s should have 5 lines, got %d, error: %v", file1, lines, err)
	}

	if lines, err := countLines(file2); err != nil || lines != 47 {
		t.Errorf("File %s should have 47 lines, got %d, error: %v", file2, lines, err)
	}

	cleanUp()

}

func countLines(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(strings.Split(string(data), "\n")), nil
}

func cleanUp() {
	filesToRemove, err := filepath.Glob("./*_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesToRemove {
		if err := os.Remove(file); err != nil {
			log.Fatal(err)
		}
	}
}
