package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	c := 52
	f := "Ace of Spades"
	l := "King of Clubs"

	if len(d) != c {
		t.Errorf("Expected deck length of %v but got %v", c, len(d))
	}

	if d[0] != f {
		t.Errorf("Expected first card to be %v but got %v", f, d[0])
	}

	if d[len(d)-1] != l {
		t.Errorf("Expected last card to be %v but got %v", l, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	c := 52

	if len(loadedDeck) != c {
		t.Errorf("Expected length of loaded deck to equal %v but got %v", c, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
