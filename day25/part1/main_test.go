package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStealSharedKey(t *testing.T) {
	cardPublicKey := uint(5764801)
	doorPublicKey := uint(17807724)
	assert.Equal(t, uint(14897079), StealSharedKey(cardPublicKey, doorPublicKey))
}

func TestSolveByBabyStepGiantStep(t *testing.T) {
	assert.Equal(t, uint(8), SolveByBabyStepGiantStep(7, 5764801, primeNumber))
	assert.Equal(t, uint(1000), SolveByBabyStepGiantStep(7, 3101882, primeNumber))
}
