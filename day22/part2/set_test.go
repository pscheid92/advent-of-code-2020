package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeckSet_Add(t *testing.T) {
	deck1 := NewDeck(1, 2, 3)
	deck2 := NewDeck(10, 2, 3)
	deck3 := NewDeck(1, 2, 3)

	d := DeckSet{}
	assert.Len(t, d, 0)

	d.Add(deck1)
	assert.Len(t, d, 1)
	assert.Contains(t, d, "1;2;3;")

	d.Add(deck2)
	assert.Len(t, d, 2)
	assert.Contains(t, d, "1;2;3;")
	assert.Contains(t, d, "10;2;3;")

	d.Add(deck3)
	assert.Len(t, d, 2)
	assert.Contains(t, d, "1;2;3;")
	assert.Contains(t, d, "10;2;3;")
}

func TestDeckSet_Contains(t *testing.T) {
	seenDeck := NewDeck(1, 2, 3)
	unseenDeck := NewDeck(1, 2, 30)

	d := DeckSet{}
	assert.False(t, d.Contains(seenDeck))
	assert.False(t, d.Contains(unseenDeck))

	d.Add(seenDeck)
	assert.True(t, d.Contains(seenDeck))
	assert.False(t, d.Contains(unseenDeck))
}
