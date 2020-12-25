package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var input string = `abc

a
b
c

ab
ac

a
a
a
a

b
`

func TestCountUniqueLetters(t *testing.T) {
	lines := helpers.ReadLineByLine(strings.NewReader(input))
	groups := helpers.StackMultilineSeparatedByEmptyOne(lines)
	expected := []int{3, 3, 3, 1, 1}

	for i, g := range groups {
		assert.Equal(t, expected[i], CountUniqueLetters(g))
	}
}
