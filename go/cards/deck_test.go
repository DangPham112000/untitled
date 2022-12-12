package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spaces but %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected last card is Four of Hearts but %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.shuffer()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length 16 but %v", len(loadedDeck))
	}

	if loadedDeck[0] != deck[0] {
		t.Errorf("Expected loaded deck is the same old but %v", loadedDeck[0])
	}
}
