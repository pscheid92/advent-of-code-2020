package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModPower(t *testing.T) {
	assert.Equal(t, uint(5764801), Power(7, 8, 20201227))
	assert.Equal(t, uint(3101882), Power(7, 1000, 20201227))
}
