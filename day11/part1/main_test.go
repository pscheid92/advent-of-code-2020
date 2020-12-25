package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateOneStep(t *testing.T) {
	lines, err := helpers.ReadLineByLineFromFile("testdata/test_simulate_one_step.txt")
	if err != nil {
		assert.FailNow(t, "error reading test data")
	}

	states := helpers.GroupMultilineSeparatedByEmptyOne(lines)
	simulator := NewSeatsSimulation(states[0])

	for _, expected := range states[1:] {
		changed := simulator.Simulate()
		assert.True(t, changed)
		assert.Equal(t, expected, simulator.SeatsAsStringsSlice())
	}
}

func TestSimulateUntilSteadyState(t *testing.T) {
	initalSeats := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	simulator := NewSeatsSimulation(initalSeats)
	steadyState := simulator.SimulateUntilSteadyState(1_000_000)
	occupiedSeats := simulator.CountOccupiedSeats()

	assert.True(t, steadyState)
	assert.Equal(t, 37, occupiedSeats)
}
