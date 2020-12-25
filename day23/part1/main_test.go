package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram(t *testing.T) {
	cases := []struct {
		input  string
		moves  int
		output string
	}{
		{"389125467", 10, "92658374"},
		{"389125467", 100, "67384529"},
	}

	for _, c := range cases {
		currentGame, min, max := CreateGame(c.input)
		currentGame = PlayNMoves(c.moves, currentGame, min, max)
		output, err := GenerateOutput(currentGame)
		assert.NoError(t, err)
		assert.Equal(t, c.output, output)
	}
}
