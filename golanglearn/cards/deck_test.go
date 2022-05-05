package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("Expected Spades of Ace buf %v", d[0])
	}
	if d[len(d)-1] != "clubs of Four" {
		t.Errorf("Expected Spades of Ace buf %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remaingCard := deal(d, 2)

	if len(remaingCard) != 14 {
		t.Errorf("Expected deck length of 14, but got %v", len(remaingCard))
	}
	if len(hand) != 2 {
		t.Errorf("Expected deck length of 2, but got %v", len(hand))
	}
}

func TestShuffle(t *testing.T) {
	b := newDeck()
	d := newDeck()
	d.shuffle()
	e := false
	for i, card := range d {
		if card != b[i] {
			e = true
		}
	}
	if e == false {
		t.Errorf("Expected deck is Shuffled")
	}
}
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingfile")

	deck := newDeck()
	deck.saveToFile("_testingfile")

	loaded := newDeckFromFile("_testingfile")
	if len(loaded) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loaded))
	}
	if loaded[0] != "Spades of Ace" {
		t.Errorf("Expected Spades of Ace buf %v", loaded[0])
	}
	if loaded[len(loaded)-1] != "clubs of Four" {
		t.Errorf("Expected Spades of Ace buf %v", loaded[len(loaded)-1])
	}

	os.Remove("_testingfile")
}
