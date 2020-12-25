package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountValidProblems(t *testing.T) {
	input := []ProblemInstance{
		{1, 3, 'a', "abcde"},
		{1, 3, 'b', "cdefg"},
		{1, 3, 'c', "ccccccccc"},
	}

	count := CountValidProblems(input)
	assert.Equal(t, 1, count)
}
