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
		{ShiftNorth, 3},
		{MoveForward, 7},
		{RotateRight, 90},
		{MoveForward, 11},
	}

	instructions, err := ExtractInstructions(lines)
	assert.NoError(t, err)
	assert.Equal(t, expected, instructions)
}

func TestShip_Creation(t *testing.T) {
	ship := Ship{}
	assert.Equal(t, 0, ship.X)
	assert.Equal(t, 0, ship.Y)
	assert.Equal(t, 0, ship.Degree)
}

func TestShip_ExecuteInstruction(t *testing.T) {
	cases := []struct {
		instruction Instruction
		state       Ship
	}{
		{Instruction{MoveForward, 10}, Ship{10, 0, 0}},
		{Instruction{ShiftNorth, 3}, Ship{10, 3, 0}},
		{Instruction{MoveForward, 7}, Ship{17, 3, 0}},
		{Instruction{RotateRight, 90}, Ship{17, 3, 270}},
		{Instruction{MoveForward, 11}, Ship{17, -8, 270}},
	}

	ship := Ship{}
	for _, c := range cases {
		ship.ExecuteInstruction(c.instruction)
		assert.Equal(t, c.state, ship)
	}
}

func TestShip_ExecuteMultipleInstructions(t *testing.T) {
	instructions := []Instruction{
		{MoveForward, 10},
		{ShiftNorth, 3},
		{MoveForward, 7},
		{RotateRight, 90},
		{MoveForward, 11},
	}

	expected := Ship{17, -8, 270}

	ship := Ship{}
	ship.ExecuteMultipleInstructions(instructions)
	assert.Equal(t, expected, ship)
}

func TestShip_GetManhattanDistance(t *testing.T) {
	ship := Ship{17, -8, 270}
	distance := ship.GetManhattanDistance()
	assert.Equal(t, 25, distance)
}

func TestNormalizeDegree(t *testing.T) {
	cases := [][2]int{
		{0, 0},
		{90, 90},
		{180, 180},
		{270, 270},
		{360, 0},
		{450, 90},
		{2610, 90},
		{-90, 270},
		{-180, 180},
		{-1080, 0},
	}

	for _, c := range cases {
		assert.Equal(t, c[1], NormalizeDegree(c[0]))
	}
}

func TestDegreeToRadiant(t *testing.T) {
	degrees := [...]int{0, 90, 180, 270, 360}
	radiants := [...]float64{0.0, math.Pi / 2, math.Pi, math.Pi * 3.0 / 2.0, 2 * math.Pi}

	for i := 0; i < len(degrees); i++ {
		assert.Equal(t, radiants[i], DegreeToRadiant(degrees[i]))
	}
}
