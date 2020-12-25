package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestExtractInstructions(t *testing.T) {
	lines := []string{"F10", "N3", "F7", "R90", "F11"}
	expected := []Instruction{
		{MoveForward, 10},
		{MoveWaypointNorth, 3},
		{MoveForward, 7},
		{RotateWaypointRight, 90},
		{MoveForward, 11},
	}

	instructions, err := ExtractInstructions(lines)
	assert.NoError(t, err)
	assert.Equal(t, expected, instructions)
}

func TestShip_Creation(t *testing.T) {
	ship := NewShip()
	assert.Equal(t, 0, ship.X)
	assert.Equal(t, 0, ship.Y)
	assert.Equal(t, 10, ship.WaypointX)
	assert.Equal(t, 1, ship.WaypointY)
}

func TestShip_ExecuteInstruction(t *testing.T) {
	cases := []struct {
		instruction Instruction
		state       Ship
	}{
		{Instruction{MoveForward, 10}, Ship{100, 10, 10, 1}},
		{Instruction{MoveWaypointNorth, 3}, Ship{100, 10, 10, 4}},
		{Instruction{MoveForward, 7}, Ship{170, 38, 10, 4}},
		{Instruction{RotateWaypointRight, 90}, Ship{170, 38, 4, -10}},
		{Instruction{MoveForward, 11}, Ship{214, -72, 4, -10}},
	}

	ship := NewShip()
	for _, c := range cases {
		ship.ExecuteInstruction(c.instruction)
		assert.Equal(t, c.state, ship)
	}
}

func TestShip_ExecuteMultipleInstructions(t *testing.T) {
	instructions := []Instruction{
		{MoveForward, 10},
		{MoveWaypointNorth, 3},
		{MoveForward, 7},
		{RotateWaypointRight, 90},
		{MoveForward, 11},
	}

	expected := Ship{214, -72, 4, -10}

	ship := NewShip()
	ship.ExecuteMultipleInstructions(instructions)
	assert.Equal(t, expected, ship)
}

func TestShip_GetManhattanDistance(t *testing.T) {
	ship := Ship{214, -72, 4, -10}
	distance := ship.GetManhattanDistance()
	assert.Equal(t, 286, distance)
}

func TestDegreeToRadiant(t *testing.T) {
	degrees := [...]int{0, 90, 180, 270, 360}
	radiants := [...]float64{0.0, math.Pi / 2, math.Pi, math.Pi * 3.0 / 2.0, 2 * math.Pi}

	for i := 0; i < len(degrees); i++ {
		assert.Equal(t, radiants[i], DegreeToRadiant(degrees[i]))
	}
}
