package main

type State rune

const (
	Inactive State = '.'
	Active   State = '#'
)

const Dimension = 3

type Coordinate [Dimension]int

type World struct {
	activePoints map[Coordinate]byte
}

func InitWorld(lines []string) World {
	world := World{activePoints: map[Coordinate]byte{}}

	for i, line := range lines {
		for j, state := range line {
			if State(state) == Active {
				coordinate := Coordinate{i, j, 0}
				world.activePoints[coordinate] = 0
			}
		}
	}

	return world
}

func (w World) CountActive() int {
	return len(w.activePoints)
}

func (w World) GatherActivePoints() []Coordinate {
	activePoints := make([]Coordinate, 0, len(w.activePoints))
	for c := range w.activePoints {
		activePoints = append(activePoints, c)
	}
	return activePoints
}

func (w World) SimulateSteps(steps int) {
	for i := 0; i < steps; i++ {
		// phase 1: visit points and collect change instructions
		changeInstructions := make(map[Coordinate]State)
		for _, point := range w.GatherInterestingPoints() {
			state := w.Get(point)
			count := w.CountActiveNeighbors(point)

			// inactive and exactly 3 active neighbors -> active
			if state == Inactive && count == 3 {
				changeInstructions[point] = Active
			}

			// active and not exactly 2 or 3 active neighbors -> inactive
			if state == Active && count != 2 && count != 3 {
				changeInstructions[point] = Inactive
			}
		}

		// phase 2: apply change instructions
		for coordinate, state := range changeInstructions {
			w.Set(coordinate, state)
		}
	}
}

func (w World) Get(coordinate Coordinate) State {
	if _, ok := w.activePoints[coordinate]; ok {
		return Active
	}
	return Inactive
}

func (w World) Set(coordinate Coordinate, state State) {
	_, ok := w.activePoints[coordinate]

	// is active but should be inactive, delete it
	if ok && state == Inactive {
		delete(w.activePoints, coordinate)
	}

	// is inactive but should be active, add it
	if !ok && state == Active {
		w.activePoints[coordinate] = 0
	}
}

func (w World) CountActiveNeighbors(middle Coordinate) (sum int) {
	for _, n := range GetNeighborCoordinates(middle) {
		if w.Get(n) == Active {
			sum += 1
		}
	}
	return sum
}

func (w World) GatherInterestingPoints() []Coordinate {
	// every active point and their neighbors are points of interest (poi).
	// every point that is more than one point afar, from any other active point cannot change. therefore, we ignore them here.
	poi := make(map[Coordinate]byte)
	for current := range w.activePoints {
		poi[current] = 0
		for _, neighbor := range GetNeighborCoordinates(current) {
			poi[neighbor] = 0
		}
	}

	// collect distinct poi coordinates from map
	distinctPoints := make([]Coordinate, 0, len(poi))
	for p := range poi {
		distinctPoints = append(distinctPoints, p)
	}
	return distinctPoints
}

func GetNeighborCoordinates(middle Coordinate) []Coordinate {
	scheme := [3]int{-1, +0, +1}

	// calculate dimensions
	dimensions := make([]int, Dimension)
	for i := 0; i < Dimension; i++ {
		dimensions[i] = len(scheme)
	}

	coordinates := make([]Coordinate, 0)
	for p := InitPermutation(dimensions...); p.Next(); {
		neighbor := middle
		val := p.Value()
		for i, x := range val {
			diff := scheme[x]
			neighbor[i] += diff
		}

		if neighbor != middle {
			coordinates = append(coordinates, neighbor)
		}
	}

	return coordinates
}
