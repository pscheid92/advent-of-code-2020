package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindEntriesSummingTo(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	result, ok := FindTreeEntriesSummingTo(input, 2020)
	assert.True(t, ok)
	assert.ElementsMatch(t, [...]int{979, 366, 675}, result)
}
