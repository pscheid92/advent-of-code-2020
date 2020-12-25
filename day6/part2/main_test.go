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

func TestCountCommonAnswers(t *testing.T) {
	expected := []int{3, 0, 1, 1, 1}

	lines := helpers.ReadLineByLine(strings.NewReader(input))
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)
	assert.Len(t, groups, 5)

	for i, g := range groups {
		assert.Equal(t, expected[i], CountCommonAnswers(g))
	}
}
