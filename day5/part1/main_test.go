package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeSeat(t *testing.T) {
	input := "FBFBBFFRLR"
	id := DecodeSeat(input)
	assert.Equal(t, id, 357)
}
