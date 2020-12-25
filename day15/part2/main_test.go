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
		{[]int{0, 3, 6}, 175594},
		{[]int{1, 3, 2}, 2578},
		{[]int{2, 1, 3}, 3544142},
		{[]int{1, 2, 3}, 261214},
		{[]int{2, 3, 1}, 6895259},
		{[]int{3, 2, 1}, 18},
		{[]int{3, 1, 2}, 362},
	}

	for _, c := range cases {
		finalNumber := PlayGame(c.startSequence, 30000000)
		assert.Equal(t, c.expectedFinalNumber, finalNumber)
	}
}
