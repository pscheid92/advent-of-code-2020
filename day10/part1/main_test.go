package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountJoltDifferences(t *testing.T) {
	cases := []struct {
		jolts  []int
		count1 int
		count3 int
	}{
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 7, 5},
		{[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 22, 10},
	}

	for _, c := range cases {
		diffs := CountJoltDifferences(c.jolts)
		assert.Equal(t, c.count1, diffs[1])
		assert.Equal(t, c.count3, diffs[3])
	}
}
