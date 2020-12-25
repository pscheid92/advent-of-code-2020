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

	ship := NewShip()
	ship.ExecuteMultipleInstructions(instructions)

	distance := ship.GetManhattanDistance()
	fmt.Printf("solution code: %d\n", distance)
}

type Command rune

const (
	MoveWaypointNorth   Command = 'N'
	MoveWaypointSouth           = 'S'
	MoveWaypointWest            = 'W'
	MoveWaypointEast            = 'E'
	RotateWaypointLeft          = 'L'
	RotateWaypointRight         = 'R'
	MoveForward                 = 'F'
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
	X         int
	Y         int
	WaypointX int
	WaypointY int
}

func NewShip() Ship {
	return Ship{
		X:         0,
		Y:         0,
		WaypointX: 10,
		WaypointY: 1,
	}
}

func (s *Ship) ExecuteMultipleInstructions(instructions []Instruction) {
	for _, i := range instructions {
		s.ExecuteInstruction(i)
	}
}

func (s *Ship) RotateWaypoint(degree int) {
	for i := 0; i < degree/90; i++ {
		x, y := s.WaypointX, s.WaypointY
		s.WaypointX = -y
		s.WaypointY = x
	}
}

func (s *Ship) ExecuteInstruction(instruction Instruction) {
	switch instruction.Command {
	case MoveWaypointNorth:
		s.WaypointY += instruction.Units
	case MoveWaypointSouth:
		s.WaypointY -= instruction.Units
	case MoveWaypointWest:
		s.WaypointX -= instruction.Units
	case MoveWaypointEast:
		s.WaypointX += instruction.Units
	case RotateWaypointLeft:
		s.RotateWaypoint(instruction.Units)
	case RotateWaypointRight:
		s.RotateWaypoint(360 - instruction.Units)
	case MoveForward:
		s.X += instruction.Units * s.WaypointX
		s.Y += instruction.Units * s.WaypointY
	default:
		panic("i should not be here!")
	}
}

func (s Ship) GetManhattanDistance() int {
	x := math.Abs(float64(s.X))
	y := math.Abs(float64(s.Y))
	return int(x + y)
}

func DegreeToRadiant(degree int) float64 {
	return float64(degree) * (math.Pi / 180.0)
}
