package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	seats, err := helpers.ReadLineByLineFromFile("day11/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	simulator := NewSeatsSimulation(seats)
	steadyState := simulator.SimulateUntilSteadyState(1_000_000)
	if !steadyState {
		log.Fatalln("hit iteration limit")
	}

	count := simulator.CountOccupiedSeats()
	fmt.Printf("solution code: %d\n", count)
}

const Floor = '.'
const EmptySeat = 'L'
const OccupiedSeat = '#'

type SeatsSimulator struct {
	Seats   [][]rune
	Rows    int
	Columns int
}

func NewSeatsSimulation(seats []string) SeatsSimulator {
	simulator := SeatsSimulator{
		Rows:    len(seats),
		Columns: len(seats[0]),
		Seats:   make([][]rune, len(seats)),
	}

	for i, row := range seats {
		simulator.Seats[i] = []rune(row)
	}

	return simulator
}

func (s SeatsSimulator) CountOccupiedSeats() int {
	counter := 0
	for _, row := range s.Seats {
		for _, seat := range row {
			if seat == OccupiedSeat {
				counter++
			}
		}
	}
	return counter
}

func (s SeatsSimulator) SeatsAsStringsSlice() []string {
	output := make([]string, s.Rows)
	for i, row := range s.Seats {
		output[i] = string(row)
	}
	return output
}

func (s SeatsSimulator) SimulateUntilSteadyState(iterationMaximum uint64) bool {
	reachedSteadyState := false

	for i := uint64(0); i < iterationMaximum; i++ {
		if changed := s.Simulate(); !changed {
			reachedSteadyState = true
			break
		}
	}

	return reachedSteadyState
}

func (s SeatsSimulator) Simulate() bool {
	changeRows := make([]int, 0)
	changeCols := make([]int, 0)
	changeStates := make([]rune, 0)

	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {

			// floor will always remain empty
			if s.Seats[i][j] == Floor {
				continue
			}

			count := s.countAdjacentSeats(i, j)

			if s.Seats[i][j] == EmptySeat && count == 0 {
				changeRows = append(changeRows, i)
				changeCols = append(changeCols, j)
				changeStates = append(changeStates, OccupiedSeat)
				continue
			}

			if s.Seats[i][j] == OccupiedSeat && count >= 4 {
				changeRows = append(changeRows, i)
				changeCols = append(changeCols, j)
				changeStates = append(changeStates, EmptySeat)
				continue
			}
		}
	}

	for i := 0; i < len(changeStates); i++ {
		x := changeRows[i]
		y := changeCols[i]
		r := changeStates[i]
		s.Seats[x][y] = r
	}

	return len(changeStates) > 0
}

func (s SeatsSimulator) countAdjacentSeats(row, column int) int {
	counter := 0

	stencil := [][]int{
		{-1, -1},
		{-1, +0},
		{-1, +1},
		{+0, -1},
		{+0, +1},
		{+1, -1},
		{+1, +0},
		{+1, +1},
	}

	for _, diff := range stencil {
		rd, cd := diff[0], diff[1]

		// upper line edge detection
		if row == 0 && rd == -1 {
			continue
		}

		// lower line edge detection
		if row == s.Rows-1 && rd == +1 {
			continue
		}

		// left side edge detection
		if column == 0 && cd == -1 {
			continue
		}

		// right side edge detection
		if column == s.Columns-1 && cd == +1 {
			continue
		}

		// actual counting
		if s.Seats[row+rd][column+cd] == OccupiedSeat {
			counter++
		}
	}

	return counter
}
