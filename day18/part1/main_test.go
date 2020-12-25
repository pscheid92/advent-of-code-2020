package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluate(t *testing.T) {
	cases := []struct {
		expression string
		result     int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, c := range cases {
		result, err := Evaluate(c.expression)
		assert.NoError(t, err, "received error evaluating '%s': %s", c.expression, err)
		assert.Equal(t, c.result, result, "evaluated to %d, expected %d for expression '%s'", result, c.result, c.expression)
	}
}
