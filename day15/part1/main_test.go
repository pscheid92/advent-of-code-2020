package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		startSequence       []int
		expectedFinalNumber int
	}{
		{[]int{0, 3, 6}, 436},
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}

	for _, c := range cases {
		finalNumber := PlayGame(c.startSequence, 2020)
		assert.Equal(t, c.expectedFinalNumber, finalNumber)
	}
}
