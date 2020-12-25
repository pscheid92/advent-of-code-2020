package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSide_Opposite(t *testing.T) {
	cases := [][2]Side{
		{UpperSide, LowerSide},
		{RightSide, LeftSide},
		{LowerSide, UpperSide},
		{LeftSide, RightSide},
	}

	for _, c := range cases {
		assert.Equal(t, c[1], c[0].Opposite())
	}
}

func TestNewTile(t *testing.T) {
	tile := NewTile(123, Flipped, RotatedLeft, []string{"12", "34"})
	assert.Equal(t, 123, tile.ID)
	assert.Equal(t, []string{"12", "34"}, tile.data)
	assert.Equal(t, Flipped, tile.Orientation)
	assert.Equal(t, RotatedLeft, tile.Rotation)
	assert.Equal(t, [4]string{"12", "24", "34", "13"}, tile.borders)

	for _, x := range tile.neighbors {
		assert.Equal(t, 0, x)
	}
}

func TestReadTile(t *testing.T) {
	block := []string{
		"Tile 2311:",
		"..##.#..#.",
		"##..#.....",
		"#...##..#.",
		"####.#...#",
		"##.##.###.",
		"##...#.###",
		".#.#.#..##",
		"..#....#..",
		"###...#.#.",
		"..###..###",
	}

	tile, err := ReadTile(block)
	assert.NoError(t, err)
	assert.Equal(t, 2311, tile.ID)
	assert.Equal(t, block[1:], tile.data)
	assert.Equal(t, Original, tile.Orientation)
	assert.Equal(t, RotatedUp, tile.Rotation)
}

func TestTile_Flip(t *testing.T) {
	inputTile := NewTile(0, Original, RotatedUp, []string{"123", "456", "789"})
	expectedTile := NewTile(0, Flipped, RotatedUp, []string{"321", "654", "987"})

	// flipping once should result in expectedTile
	onceFlipped := inputTile.Flip()
	assert.Equal(t, expectedTile, onceFlipped)

	// flipping again should result in the original inputTile
	twiceFlipped := onceFlipped.Flip()
	assert.Equal(t, inputTile, twiceFlipped)
}

func TestTile_Rotate(t *testing.T) {
	inputTile := NewTile(0, Original, RotatedUp, []string{"12", "34"})
	expectedRotatedOnce := NewTile(0, Original, RotatedRight, []string{"31", "42"})
	expectedRotatedTwice := NewTile(0, Original, RotatedDown, []string{"43", "21"})
	expectedRotatedThrice := NewTile(0, Original, RotatedLeft, []string{"24", "13"})

	// rotating first time (showing to the right)
	rotatedOnce := inputTile.Rotate()
	assert.Equal(t, expectedRotatedOnce, rotatedOnce)

	// rotating second time (upside down)
	rotatedTwice := rotatedOnce.Rotate()
	assert.Equal(t, expectedRotatedTwice, rotatedTwice)

	// rotating third time (showing to the left)
	rotatedThrice := rotatedTwice.Rotate()
	assert.Equal(t, expectedRotatedThrice, rotatedThrice)

	// rotating again should lead to original input tile
	rotatedQuadrice := rotatedThrice.Rotate()
	assert.Equal(t, inputTile, rotatedQuadrice)
}

func TestTile_GetBorder(t *testing.T) {
	tile := NewTile(0, Original, RotatedUp, []string{"12", "34"})
	assert.Equal(t, "12", tile.GetBorder(UpperSide))
	assert.Equal(t, "24", tile.GetBorder(RightSide))
	assert.Equal(t, "34", tile.GetBorder(LowerSide))
	assert.Equal(t, "13", tile.GetBorder(LeftSide))
}

func TestTile_GetNeighbor(t *testing.T) {
	data := []string{"xx", "xx"}
	tile := NewTile(99, Original, RotatedUp, data)

	// check unconnected sides
	for i := 0; i < 4; i++ {
		assert.Equal(t, 0, tile.neighbors[i])
	}

	// connect the sides
	for i := 0; i < 4; i++ {
		neighbor := NewTile(i, Original, RotatedUp, data)
		tile.neighbors[i] = neighbor.ID
	}

	// check again
	for i := 0; i < 4; i++ {
		assert.Equal(t, i, tile.GetNeighbor(Side(i)))
	}
}

func TestTile_SetNeighbor(t *testing.T) {
	// prepare a tile and its future neighbor
	data := []string{"xx", "xx"}
	tile := NewTile(123, Original, RotatedUp, data)
	other := NewTile(321, Flipped, RotatedRight, data)

	// couple them
	tile.SetNeighbor(LeftSide, &other)

	// check on tile
	assert.NotNil(t, tile.neighbors[LeftSide])
	assert.Equal(t, other.ID, tile.GetNeighbor(LeftSide))

	// check on neighbor
	assert.NotNil(t, other.neighbors[RightSide])
	assert.Equal(t, tile.ID, other.GetNeighbor(RightSide))
}

func TestTile_CountNeighbors(t *testing.T) {
	// prepare a tile and its future neighbor
	data := []string{"xx", "xx"}
	tile := NewTile(123, Original, RotatedUp, data)
	other := NewTile(321, Flipped, RotatedRight, data)

	// counting should return 0 for both
	assert.Equal(t, 0, tile.CountNeighbors())
	assert.Equal(t, 0, other.CountNeighbors())

	// couple them
	tile.SetNeighbor(LeftSide, &other)

	// now both should have one neighbor
	assert.Equal(t, 1, tile.CountNeighbors())
	assert.Equal(t, 1, other.CountNeighbors())
}

func TestGatherBorder(t *testing.T) {
	data := []string{"123", "456", "789"}

	cases := []struct {
		side     Side
		expected string
	}{
		{UpperSide, "123"},
		{LowerSide, "789"},
		{LeftSide, "147"},
		{RightSide, "369"},
	}

	for _, c := range cases {
		border := GatherBorder(data, c.side)
		assert.Equal(t, c.expected, border)
	}
}

func TestReverseString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"X", "X"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"neffen", "neffen"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, ReverseString(c.input))
	}
}

func TestMatchTiles(t *testing.T) {
	tile := NewTile(1, Original, RotatedLeft, []string{"12", "34"})
	matchingTile := NewTile(2, Original, RotatedUp, []string{"2X", "4X"})
	unmatchingTile := NewTile(3, Flipped, RotatedDown, []string{"##", "##"})

	// never matches with itself
	assert.False(t, MatchTiles(&tile, &tile))

	// does not match works
	assert.False(t, MatchTiles(&tile, &unmatchingTile))
	assert.Equal(t, 0, tile.CountNeighbors())

	// matching does work (on both tiles)
	assert.True(t, MatchTiles(&tile, &matchingTile))
	assert.Equal(t, 1, tile.CountNeighbors())
	assert.Equal(t, 1, matchingTile.CountNeighbors())
}
