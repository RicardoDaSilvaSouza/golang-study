package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	assert := assert.New(t)

	cards := newDeck()

	assert.Lenf(cards, 16, "Length expected %d found %d", 16, len(cards))
	assert.Equalf("Ace of Spades", cards[0], "Card at index 0 was expected %s found %s", "Ace of Spades", cards[0])
	assert.Equalf("Four of Clubs", cards[len(cards)-1], "Card at index %d was expected %s found %s", len(cards)-1, "Ace of Spades", cards[0])

}
