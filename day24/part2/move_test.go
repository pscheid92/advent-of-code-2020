package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePaths(t *testing.T) {
	input := []string{"sesenwnenenewseeswwswswwnenewsewsw", "seswneswswsenwwnwse"}
	expectedPaths := [][]Move{{SE, SE, NW, NE, NE, NE, W, SE, E, SW, W, SW, SW, W, NE, NE, W, SE, W, SW}, {SE, SW, NE, SW, SW, SE, NW, W, NW, SE}}

	paths, err := ParsePaths(input)
	assert.NoError(t, err)
	assert.Equal(t, expectedPaths, paths)
}
