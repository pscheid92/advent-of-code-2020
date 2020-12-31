package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram(t *testing.T) {
	input := "389125467"
	size := 1_000_000
	moves := 10_000_000
	expectedOutput := 149245887792

	currentGame, min, max, lookup := CreateGame(input, size)
	_ = PlayNMoves(moves, currentGame, min, max, lookup)

	output, err := CalculateSolutionCode(lookup)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
