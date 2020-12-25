package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateModularInverse(t *testing.T) {
	cases := [][3]uint64{
		{56, 5, 1},
		{40, 7, 3},
		{35, 8, 3},
	}

	for _, c := range cases {
		inverse := CalculateModularInverse(c[0], c[1])
		assert.Equal(t, c[2], inverse)
	}
}

func TestChineseReminderTheorem(t *testing.T) {
	cases := []struct {
		line   string
		expect uint64
	}{
		{"7,13,x,x,59,x,31,19", uint64(1068781)},
		{"17,x,13,19", uint64(3417)},
		{"67,7,59,61", uint64(754018)},
		{"67,x,7,59,61", uint64(779210)},
		{"67,7,x,59,61", uint64(1261476)},
		{"1789,37,47,1889", uint64(1202161486)},
	}

	for i, c := range cases {
		busses, err := ParseInput(c.line)
		assert.NoError(t, err)

		ns, bs := PrepareInputForCRT(busses)
		assert.Equal(t, c.expect, ChineseReminderTheorem(ns, bs), "error in case %d", i)
	}
}
