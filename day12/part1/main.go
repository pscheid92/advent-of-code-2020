package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"math"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day12/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	instructions, err := ExtractInstructions(lines)
	if err != nil {
		log.Fatalln(err)
	}

	ship := Ship{}
	ship.ExecuteMultipleInstructions(instructions)

	distance := ship.GetManhattanDistance()
	fmt.Printf("solution code: %d\n", distance)
}

type Command rune

const (
	ShiftNorth  Command = 'N'
	ShiftSouth          = 'S'
	ShiftWest           = 'W'
	ShiftEast           = 'E'
	RotateLeft          = 'L'
	RotateRight         = 'R'
	MoveForward         = 'F'
)

type Instruction struct {
	Command Command
	Units   int
}

func ExtractInstructions(lines []string) ([]Instruction, error) {
	instructions := make([]Instruction, len(lines))
	for i, l := range lines {
		n, err := fmt.Sscanf(l, "%c%d", &instructions[i].Command, &instructions[i].Units)
		if err != nil {
			return nil, fmt.Errorf("error parsing '%s': %w", l, err)
		}
		if n != 2 {
			return nil, fmt.Errorf("expected 2, found %d values in '%s'", n, l)
		}
	}
	return instructions, nil
}

type Ship struct {
	X      int
	Y      int
	Degree int
}

func (s *Ship) ExecuteMultipleInstructions(instructions []Instruction) {
	for _, i := range instructions {
		s.ExecuteInstruction(i)
	}
}

func (s *Ship) ExecuteInstruction(instruction Instruction) {
	switch instruction.Command {
	case ShiftNorth:
		s.Y += instruction.Units
	case ShiftSouth:
		s.Y -= instruction.Units
	case ShiftWest:
		s.X -= instruction.Units
	case ShiftEast:
		s.X += instruction.Units
	case RotateLeft:
		s.Degree = NormalizeDegree(s.Degree + instruction.Units)
	case RotateRight:
		s.Degree = NormalizeDegree(s.Degree - instruction.Units)
	case MoveForward:
		units := float64(instruction.Units)
		radiant := DegreeToRadiant(s.Degree)
		s.X += int(units * math.Cos(radiant))
		s.Y += int(units * math.Sin(radiant))
	default:
		panic("i should not be here!")
	}
}

func (s Ship) GetManhattanDistance() int {
	x := math.Abs(float64(s.X))
	y := math.Abs(float64(s.Y))
	return int(x + y)
}

func NormalizeDegree(degree int) int {
	degree %= 360
	if degree < 0 {
		degree += 360
	}
	return degree
}

func DegreeToRadiant(degree int) float64 {
	return float64(degree) * (math.Pi / 180.0)
}
