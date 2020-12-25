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
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, c := range cases {
		result, err := Evaluate(c.expression)
		assert.NoError(t, err, "received error evaluating '%s': %s", c.expression, err)
		assert.Equal(t, c.result, result, "evaluated to %d, expected %d for expression '%s'", result, c.result, c.expression)
	}
}
