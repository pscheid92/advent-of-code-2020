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

type VisibleNextSeats [][2]int

type SeatsSimulator struct {
	Seats            [][]rune
	Rows             int
	Columns          int
	VisibleNextSeats [][]VisibleNextSeats
}

func NewSeatsSimulation(seats []string) SeatsSimulator {
	rows := len(seats)
	cols := len(seats[0])

	simulator := SeatsSimulator{
		Rows:             rows,
		Columns:          cols,
		Seats:            make([][]rune, rows),
		VisibleNextSeats: make([][]VisibleNextSeats, rows),
	}

	for i, row := range seats {
		simulator.Seats[i] = []rune(row)
		simulator.VisibleNextSeats[i] = make([]VisibleNextSeats, cols)
	}

	simulator.initVisibleNextSeats()
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
	type change struct {
		x     int
		y     int
		state rune
	}
	plannedChanges := make([]change, 0)

	// phase 1: inspect each seat and gather planned changes
	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {
			// floor will always remain empty, so skip
			if s.Seats[i][j] == Floor {
				continue
			}

			count := s.countAdjacentSeats(i, j)

			if s.Seats[i][j] == EmptySeat && count == 0 {
				plannedChanges = append(plannedChanges, change{i, j, OccupiedSeat})
				continue
			}

			if s.Seats[i][j] == OccupiedSeat && count >= 5 {
				plannedChanges = append(plannedChanges, change{i, j, EmptySeat})
				continue
			}
		}
	}

	// phase2: apply planned changes
	for _, c := range plannedChanges {
		s.Seats[c.x][c.y] = c.state
	}
	return len(plannedChanges) > 0
}

func (s SeatsSimulator) initVisibleNextSeats() {
	directions := [8][2]int{
		{-1, +0}, // north
		{+1, +0}, // south
		{+0, -1}, // west
		{+0, +1}, // east
		{-1, -1}, // north west
		{-1, +1}, // north east
		{+1, -1}, // south west
		{+1, +1}, // south east
	}

	// broken out in lambda to make loop over seats more readable
	initParticularSeat := func(row, column int) {
		visibleNextSeats := make(VisibleNextSeats, 0)
		for _, d := range directions {
			currentRow := row
			currentCol := column

			for {
				// go one step in given direction
				currentRow += d[0]
				currentCol += d[1]

				// edge detection
				if currentRow < 0 || currentRow >= s.Rows || currentCol < 0 || currentCol >= s.Columns {
					break
				}

				// check if we found a seat
				if s.Seats[currentRow][currentCol] != Floor {
					visibleNextSeats = append(visibleNextSeats, [2]int{currentRow, currentCol})
					break
				}
			}
		}

		s.VisibleNextSeats[row][column] = visibleNextSeats
	}

	// run over all seats and init
	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {
			initParticularSeat(i, j)
		}
	}
}

func (s SeatsSimulator) countAdjacentSeats(row int, column int) int {
	counter := 0
	for _, seat := range s.VisibleNextSeats[row][column] {
		if s.Seats[seat[0]][seat[1]] == OccupiedSeat {
			counter++
		}
	}
	return counter
}
