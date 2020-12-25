package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkThroughForest(t *testing.T) {
	var forest = []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	counters := WalkThroughMultipleForests(slopes, forest)
	solution := helpers.Multiply(counters)
	assert.ElementsMatch(t, []int{2, 7, 3, 4, 2}, counters)
	assert.Equal(t, 336, solution)
}
