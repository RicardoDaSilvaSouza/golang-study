package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const file_path = "_decktest.txt"

func TestNewDeck(t *testing.T) {
	assert := assert.New(t)

	cards := newDeck()

	assert.Lenf(cards, 16, "Length expected 16 found %d", len(cards))
	assert.Equalf("Ace of Spades", cards[0], "Card at index 0 was expected Ace of Spades found %s", cards[0])
	assert.Equalf("Four of Clubs", cards[len(cards)-1], "Card at index %d was expected Four of Clubs found %s", len(cards)-1, cards[0])
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	assert := assert.New(t)

	os.Remove(file_path)
	deck := newDeck()
	deck.saveToFile(file_path)

	loaded := newDeckFromFile(file_path)

	assert.Lenf(loaded, 16, "Length expected 16 found %d", len(loaded))

	assert.Equalf(deck, loaded, "Loaded deck isn't correct, expected [%s] found [%s]", deck, loaded)

	os.Remove(file_path)
}
