package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedCards := Deck{1, 2, 3, 4}
	cards := []int{1, 2, 3, 4}
	deck := NewDeck(cards...)
	assert.Equal(t, expectedCards, deck)
}

func TestDeck_Draw(t *testing.T) {
	deck := NewDeck(1, 2, 3)
	assert.Equal(t, 3, deck.Draw())
	assert.Equal(t, 2, deck.Draw())
	assert.Equal(t, 1, deck.Draw())
}

func TestDeck_InsertAtBottom(t *testing.T) {
	deck := NewDeck(10, 11, 12)

	// insert a single card
	deck.InsertAtBottom(0)
	assert.Equal(t, Deck{0, 10, 11, 12}, deck)

	deck.InsertAtBottom(1, 2, 3)
	assert.Equal(t, Deck{3, 2, 1, 0, 10, 11, 12}, deck)
}

func TestDeck_Empty(t *testing.T) {
	deck := NewDeck(10, 11)
	assert.False(t, deck.Empty())

	deck.Draw()
	deck.Draw()
	assert.True(t, deck.Empty())
}

func TestDeck_NotEmpty(t *testing.T) {
	deck := NewDeck(10, 11)
	assert.True(t, deck.NotEmpty())

	deck.Draw()
	deck.Draw()
	assert.False(t, deck.NotEmpty())
}

func TestDeck_Score(t *testing.T) {
	deck := NewDeck(1, 7, 4, 9, 5, 8, 6, 10, 2, 3)
	assert.Equal(t, 306, deck.Score())
}
