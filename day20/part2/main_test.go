package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariations(t *testing.T) {
	type positionalData struct {
		Orientation Orientation
		Rotation    Rotation
	}

	expectedPositionalData := []positionalData{
		{Original, RotatedUp},
		{Original, RotatedRight},
		{Original, RotatedDown},
		{Original, RotatedLeft},
		{Flipped, RotatedUp},
		{Flipped, RotatedRight},
		{Flipped, RotatedDown},
		{Flipped, RotatedLeft},
	}

	// no tiles, no variations
	assert.Len(t, Variations([]Tile{}), 0)

	// calculate variations
	original := NewTile(1, Original, RotatedUp, []string{"xx", "xx"})
	variations := Variations([]Tile{original})

	// we expect only one entry (original id)
	assert.Len(t, variations, 8)

	// gather positional Data from variations
	actual := make([]positionalData, 0)
	for _, tile := range variations {
		actual = append(actual, positionalData{tile.Orientation, tile.Rotation})
	}

	// compare expected and actual positional Data
	assert.ElementsMatch(t, expectedPositionalData, actual)
}

func TestConnectTiles(t *testing.T) {
	expected := map[int][4]int{
		2971: {0, 1489, 2729, 0},
		1489: {0, 1171, 1427, 2971},
		1171: {0, 0, 2473, 1489},
		2729: {2971, 1427, 1951, 0},
		1427: {1489, 2473, 2311, 2729},
		2473: {1171, 0, 3079, 1427},
		1951: {2729, 2311, 0, 0},
		2311: {1427, 3079, 0, 1951},
		3079: {2473, 0, 0, 2311},
	}

	lines, err := helpers.ReadLineByLineFromFile("testdata/tiles.txt")
	assert.NoError(t, err)
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)

	tiles, err := ReadTiles(groups)
	assert.NoError(t, err)

	for _, tile := range ConnectTiles(tiles) {
		neighbors, ok := expected[tile.ID]
		assert.True(t, ok)
		assert.Equal(t, neighbors, tile.neighbors)
	}
}
