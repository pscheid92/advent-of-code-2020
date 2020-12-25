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

func TestDeck_Copy(t *testing.T) {
	original := NewDeck(1, 2, 3)
	copied := original.Copy()
	assert.Equal(t, original, copied)

	// check memory independence
	copied.InsertAtBottom(5)
	assert.NotEqual(t, original, copied)
}

func TestDeck_CutCards(t *testing.T) {
	original := NewDeck(0, 1, 2, 3, 4, 5, 6, 7)
	assert.Len(t, original, 8)

	cutted := original.CutCards(3)
	assert.Len(t, cutted, 3)
	assert.Equal(t, Deck{5, 6, 7}, cutted)
}

func TestDeck_Size(t *testing.T) {
	deck := NewDeck(1, 2, 3)
	assert.Equal(t, 3, deck.Size())

	deck.Draw()
	assert.Equal(t, 2, deck.Size())

	deck.InsertAtBottom(10, 20, 30)
	assert.Equal(t, 5, deck.Size())
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
