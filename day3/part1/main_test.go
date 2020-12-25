package main

import (
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

	count := WalkThroughForest(3, 1, forest)
	assert.Equal(t, 7, count)
}
