package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLines = []string{
	".#.",
	"..#",
	"###",
}

func TestRunSimulation(t *testing.T) {
	count := RunSimulation(testLines, 6)
	assert.Equal(t, 848, count)
}
