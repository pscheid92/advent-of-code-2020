package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var initWorld = []string{
	".#.",
	"..#",
	"###",
}

func TestInitWorld(t *testing.T) {
	expectedActivePoints := []Coordinate{
		{0, 1, 0},
		{1, 2, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 0},
	}

	// create world and collect used coordinates
	world := InitWorld(initWorld)
	actualActivePoints := world.GatherActivePoints()

	// match expected actualActivePoints with ones created in world
	assert.ElementsMatch(t, expectedActivePoints, actualActivePoints)
}

func TestWorld_CountActive(t *testing.T) {
	world := InitWorld(initWorld)
	count := world.CountActive()
	assert.Equal(t, 5, count)
}

func TestWorld_GatherActivePoints(t *testing.T) {
	expectedActivePoints := []Coordinate{
		{0, 1, 0},
		{1, 2, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 0},
	}

	world := InitWorld(initWorld)
	actualActivePoints := world.GatherActivePoints()
	assert.ElementsMatch(t, expectedActivePoints, actualActivePoints)
}

func TestWorld_SimulateSteps(t *testing.T) {
	world := InitWorld(initWorld)
	world.SimulateSteps(6)
	count := world.CountActive()
	assert.Equal(t, 112, count)
}

func TestWorld_Get(t *testing.T) {
	world := InitWorld(initWorld)

	assert.Equal(t, Active, world.Get(Coordinate{0, 1, 0}))
	assert.Equal(t, Inactive, world.Get(Coordinate{0, 0, 0}))
}

func TestWorld_Set(t *testing.T) {
	// create world and check pre-test state
	world := InitWorld(initWorld)
	assert.Equal(t, Inactive, world.Get(Coordinate{0, 0, 0}))

	// set active and check again
	world.Set(Coordinate{0, 0, 0}, Active)
	assert.Equal(t, Active, world.Get(Coordinate{0, 0, 0}))

	// set inactive again and check
	world.Set(Coordinate{0, 0, 0}, Inactive)
	assert.Equal(t, Inactive, world.Get(Coordinate{0, 0, 0}))
}

func TestWorld_CountActiveNeighbors(t *testing.T) {
	cases := []struct {
		coordinate Coordinate
		count      int
	}{
		{Coordinate{0, 0, 0}, 1},
		{Coordinate{1, 1, 0}, 5},
		{Coordinate{2, 2, 1}, 3},
	}

	world := InitWorld(initWorld)
	for _, c := range cases {
		assert.Equal(t, c.count, world.CountActiveNeighbors(c.coordinate))
	}
}

func TestWorld_GatherInterestingPoints(t *testing.T) {
	expectedPOI := []Coordinate{
		// z: -1
		{0, 0, -1},
		{0, 1, -1},
		{0, 2, -1},
		{1, 0, -1},
		{1, 1, -1},
		{1, 2, -1},
		{2, 0, -1},
		{2, 1, -1},
		{2, 2, -1},

		// z: 0
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{1, 0, 0},
		{1, 1, 0},
		{1, 2, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 0},

		// z: 1
		{0, 0, 1},
		{0, 1, 1},
		{0, 2, 1},
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 1},
		{2, 0, 1},
		{2, 1, 1},
		{2, 2, 1},
	}
	world := InitWorld([]string{"...", ".#.", "..."})
	poi := world.GatherInterestingPoints()
	assert.ElementsMatch(t, expectedPOI, poi)
}

func TestGetNeighborCoordinates(t *testing.T) {
	middleCoordinate := Coordinate{0, 0, 0}
	otherCoordinate := Coordinate{4, 5, 6}

	expectedMiddleNeighbors := []Coordinate{
		// z = -1
		{-1, -1, -1},
		{-1, +0, -1},
		{-1, +1, -1},
		{+0, -1, -1},
		{+0, +0, -1},
		{+0, +1, -1},
		{+1, -1, -1},
		{+1, +0, -1},
		{+1, +1, -1},

		// z = 0
		{-1, -1, +0},
		{-1, +0, +0},
		{-1, +1, +0},
		{+0, -1, +0},
		{+0, +1, +0},
		{+1, -1, +0},
		{+1, +0, +0},
		{+1, +1, +0},

		// z = +1
		{-1, -1, +1},
		{-1, +0, +1},
		{-1, +1, +1},
		{+0, -1, +1},
		{+0, +0, +1},
		{+0, +1, +1},
		{+1, -1, +1},
		{+1, +0, +1},
		{+1, +1, +1},
	}
	expectedOtherNeighbors := []Coordinate{
		// z = 5
		{3, 4, 5},
		{3, 5, 5},
		{3, 6, 5},
		{4, 4, 5},
		{4, 5, 5},
		{4, 6, 5},
		{5, 4, 5},
		{5, 5, 5},
		{5, 6, 5},

		// z = 6
		{3, 4, 6},
		{3, 5, 6},
		{3, 6, 6},
		{4, 4, 6},
		{4, 6, 6},
		{5, 4, 6},
		{5, 5, 6},
		{5, 6, 6},

		// z = 7
		{3, 4, 7},
		{3, 5, 7},
		{3, 6, 7},
		{4, 4, 7},
		{4, 5, 7},
		{4, 6, 7},
		{5, 4, 7},
		{5, 5, 7},
		{5, 6, 7},
	}

	assert.ElementsMatch(t, expectedMiddleNeighbors, GetNeighborCoordinates(middleCoordinate))
	assert.ElementsMatch(t, expectedOtherNeighbors, GetNeighborCoordinates(otherCoordinate))
}
