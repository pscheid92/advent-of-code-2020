package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitPermutation(t *testing.T) {
	p := InitPermutation(1, 2, 3)
	assert.Equal(t, 3, p.N)
	assert.Equal(t, []int{1, 2, 3}, p.Dimensions)
	assert.Equal(t, []int{0, 0, 0}, p.state)
}

func TestPermutation_Next(t *testing.T) {
	cases := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
	}

	i := 0
	for p := InitPermutation(3, 2); p.Next(); {
		assert.Equal(t, cases[i], p.Value())
		i++
	}

	if i < len(cases) {
		assert.Fail(t, "permutation aborted before all cases were tested")
	}
}

func TestPermutation_Value(t *testing.T) {
	p := InitPermutation(2, 2)
	assert.Equal(t, []int{0, 0}, p.Value())
}

func TestPermutation_Reset(t *testing.T) {
	p := InitPermutation(2, 2, 3)

	// overwrite state (and check that it is really set)
	p.state = []int{1, 1, 2}
	assert.Equal(t, []int{1, 1, 2}, p.state)

	// reset again and check for zero state
	p.Reset()
	assert.Equal(t, []int{0, 0, 0}, p.state)
}
