package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchTerminatingProgram(t *testing.T) {
	code := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)

	acc, err := SearchTerminatingProgram(program)
	assert.NoError(t, err)
	assert.Equal(t, 8, acc)
}
