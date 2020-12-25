package main

import (
	"errors"
	"fmt"
	"strings"
)

type Orientation int

const (
	Original Orientation = iota
	Flipped
)

func (o Orientation) String() string {
	switch o {
	case Original:
		return "ORIGINAL"
	case Flipped:
		return "FLIPPED"
	default:
		return "<unknown>"
	}
}

type Rotation int

const (
	RotatedUp Rotation = iota
	RotatedRight
	RotatedDown
	RotatedLeft
)

func (r Rotation) String() string {
	switch r {
	case RotatedUp:
		return "UP"
	case RotatedRight:
		return "RIGHT"
	case RotatedDown:
		return "DOWN"
	case RotatedLeft:
		return "LEFT"
	default:
		return "<unknown>"
	}
}

type Side int

const (
	UpperSide Side = iota
	RightSide
	LowerSide
	LeftSide
)

func (s Side) Opposite() Side {
	return (s + 2) % 4
}

func (s Side) String() string {
	switch s {
	case UpperSide:
		return "UPPER"
	case RightSide:
		return "RIGHT"
	case LowerSide:
		return "LOWER"
	case LeftSide:
		return "LEFT"
	default:
		return "<unknown>"
	}
}

type Tile struct {
	ID          int
	Orientation Orientation
	Rotation    Rotation
	Data        []string

	borders   [4]string
	neighbors [4]int
}

func NewTile(id int, orientation Orientation, rotation Rotation, data []string) Tile {
	var borders [4]string
	for i := 0; i < 4; i++ {
		borders[i] = GatherBorder(data, Side(i))
	}

	return Tile{
		ID:          id,
		Orientation: orientation,
		Rotation:    rotation,
		Data:        data,
		borders:     borders,
	}
}

func (t Tile) Flip() Tile {
	n := len(t.Data)
	data := make([]string, n)

	// flip Data
	for i, s := range t.Data {
		data[i] = ReverseString(s)
	}

	// recalculate borders
	var borders [4]string
	for i := 0; i < 4; i++ {
		borders[i] = GatherBorder(data, Side(i))
	}

	return Tile{
		ID:          t.ID,
		Orientation: (t.Orientation + 1) % 2,
		Rotation:    t.Rotation,
		Data:        data,
		borders:     borders,
	}
}

func (t Tile) Rotate() Tile {
	n := len(t.Data)
	data := make([]string, n)

	// fill new pixel grid by using rotation rule
	// new(i, j) <- old(n-j-1, i)
	var builder strings.Builder
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pixel := rune(t.Data[n-j-1][i])
			builder.WriteRune(pixel)
		}
		data[i] = builder.String()
		builder.Reset()
	}

	// recalculate borders
	var borders [4]string
	for i := 0; i < 4; i++ {
		borders[i] = GatherBorder(data, Side(i))
	}

	return Tile{
		ID:          t.ID,
		Orientation: t.Orientation,
		Rotation:    (t.Rotation + 1) % 4,
		Data:        data,
		borders:     borders,
	}
}

func (t Tile) GetBorder(side Side) string {
	return t.borders[side]
}

func (t Tile) GetNeighbor(side Side) int {
	return t.neighbors[side]
}

func (t *Tile) SetNeighbor(side Side, other *Tile) {
	t.neighbors[side] = other.ID
	other.neighbors[side.Opposite()] = t.ID
}

func (t Tile) CountNeighbors() int {
	counter := 0
	for _, x := range t.neighbors {
		if x != 0 {
			counter++
		}
	}
	return counter
}

func (t *Tile) RemoveBorder() {
	n := len(t.Data)
	newData := make([]string, 0, n-2)

	for i := 1; i < n-1; i++ {
		var builder strings.Builder
		for j := 1; j < n-1; j++ {
			builder.WriteRune(rune(t.Data[i][j]))
		}
		newData = append(newData, builder.String())
	}

	var borders [4]string
	for i := 0; i < 4; i++ {
		borders[i] = GatherBorder(newData, Side(i))
	}

	t.Data = newData
	t.borders = borders
}

func MatchTiles(x, y *Tile) bool {
	// never match same tile with itself
	if x.ID == y.ID {
		return false
	}

	// check all for sides
	for side := Side(0); side < 4; side++ {
		// compare sides, skip rest if no match
		if x.GetBorder(side) != y.GetBorder(side.Opposite()) {
			continue
		}

		x.SetNeighbor(side, y)
		return true
	}

	// found nothing
	return false
}

func ReadTiles(groups [][]string) ([]Tile, error) {
	tiles := make([]Tile, len(groups))
	for i, block := range groups {
		if tile, err := ReadTile(block); err != nil {
			return nil, err
		} else {
			tiles[i] = tile
		}
	}
	return tiles, nil
}

func ReadTile(block []string) (Tile, error) {
	// extract id
	var id int
	n, err := fmt.Sscanf(block[0], "Tile %d:", &id)
	if err != nil {
		return Tile{}, err
	}
	if n != 1 {
		return Tile{}, errors.New("error reading tile id")
	}

	// extract image Data
	data := make([]string, len(block)-1)
	for i, l := range block[1:] {
		data[i] = l
	}

	return NewTile(id, Original, RotatedUp, data), nil
}

func GatherBorder(data []string, side Side) string {
	var border string

	switch side {
	case UpperSide:
		border = data[0]
	case LowerSide:
		border = data[len(data)-1]
	case LeftSide:
		var builder strings.Builder
		for _, x := range data {
			builder.WriteRune(rune(x[0]))
		}
		border = builder.String()
	case RightSide:
		var builder strings.Builder
		i := len(data[0]) - 1
		for _, x := range data {
			builder.WriteRune(rune(x[i]))
		}
		border = builder.String()
	}

	return border
}

func ReverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for i, x := range s {
		runes[n-i-1] = x
	}
	return string(runes)
}
