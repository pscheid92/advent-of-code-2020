package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNextBus(t *testing.T) {
	now := 939
	busses := []int{7, 13, 59, 31, 19}

	arrivalTime, nextBus := CalculateNextBus(now, busses)
	assert.Equal(t, 5, arrivalTime)
	assert.Equal(t, 59, nextBus)
}
