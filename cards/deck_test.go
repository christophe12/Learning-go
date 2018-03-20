package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected length of 9, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be Ace of Spades, got %v", d[0])
	}

	lastElement := d[len(d)-1]

	if lastElement != "Three of Hearts" {
		t.Errorf("Expected last element to be Three of Hearts, got %v", lastElement)
	}
}

func TestSaveToFileAndTestNewDeckFromFile(t *testing.T) {
	testingFileName := "_decktesting"

	// cleaning for past tests
	os.Remove(testingFileName)

	deck := newDeck()
	deck.saveToFile(testingFileName)
	loadedDeck := newDeckFromFile(testingFileName)

	if len(loadedDeck) != 9 {
		t.Errorf("Expected deck to have a length of 9, but got %v", len(loadedDeck))
	}

	// cleaning for future tests
	os.Remove(testingFileName)
}
