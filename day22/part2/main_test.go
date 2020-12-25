package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var input = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

func TestExtractCardsFromInput(t *testing.T) {
	expectedCards := [2]Deck{
		{1, 3, 6, 2, 9},
		{10, 7, 4, 8, 5},
	}

	lines := helpers.ReadLineByLine(strings.NewReader(input))
	cards, err := ExtractCardsFromInput(lines)
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedCards, cards)
}

func TestPlayGame(t *testing.T) {
	lines := helpers.ReadLineByLine(strings.NewReader(input))
	cards, err := ExtractCardsFromInput(lines)
	assert.NoError(t, err)

	_, score := PlayGame(cards[0], cards[1])
	assert.Equal(t, 291, score)
}
