package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAtIndexValid(t *testing.T) {
	preamble := 5
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	for i := preamble; i < len(numbers); i++ {
		expect := numbers[i] != 127
		valid := IsAtIndexValid(i, preamble, numbers)
		assert.Equal(t, expect, valid)
	}
}
