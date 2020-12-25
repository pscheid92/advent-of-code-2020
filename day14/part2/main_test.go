package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateAddresses(t *testing.T) {
	cases := []struct {
		mask     string
		address  uint64
		expected []uint64
	}{
		{"000000000000000000000000000000X1001X", 42, []uint64{26, 27, 58, 59}},
		{"00000000000000000000000000000000X0XX", 26, []uint64{16, 17, 18, 19, 24, 25, 26, 27}},
	}

	for _, c := range cases {
		addresses := CalculateAddresses(c.mask, c.address)
		assert.ElementsMatch(t, c.expected, addresses)
	}
}

func TestRunProgram(t *testing.T) {
	program := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	memory := make(map[uint64]uint64)
	err := RunProgram(memory, program)

	assert.NoError(t, err)
	assert.EqualValues(t, 208, SumValues(memory))
}
