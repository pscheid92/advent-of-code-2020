package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMask(t *testing.T) {
	expectedSettingMask := uint64(0b1000000)
	expectedClearingMask := ^uint64(0b10)

	mask := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	assert.Equal(t, expectedSettingMask, mask.setMask)
	assert.Equal(t, expectedClearingMask, mask.clearMask)
}

func TestMask_Apply(t *testing.T) {
	mask := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	cases := [][2]uint64{
		{11, 73},
		{101, 101},
		{0, 64},
	}

	for _, c := range cases {
		assert.Equal(t, c[1], mask.Apply(c[0]))
	}
}

func TestMemory(t *testing.T) {
	memory := NewMemory()
	memory.Mask = NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	cases := [][2]uint64{
		{8, 11},
		{7, 101},
		{8, 0},
	}

	for _, c := range cases {
		memory.Set(c[0], c[1])
	}

	assert.EqualValues(t, 165, memory.SumData())
}
