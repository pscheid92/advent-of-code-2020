package main

import (
	"fmt"
)

type Coordinate [2]int

var movesToNeighbors = []Move{NE, E, SE, NW, W, SW}

func NewCoordinate() Coordinate {
	return Coordinate{0, 0}
}

func (c *Coordinate) Move(move Move) {
	switch move {
	case W:
		c[0]--
	case E:
		c[0]++
	case NW:
		c[1]--
	case NE:
		c[0]++
		c[1]--
	case SW:
		c[0]--
		c[1]++
	case SE:
		c[1]++
	default:
		panic("undefined move direction useds")
	}
}

func (c *Coordinate) MovePath(path []Move) {
	for _, p := range path {
		c.Move(p)
	}
}

func (c Coordinate) Copy() Coordinate {
	return Coordinate{c[0], c[1]}
}

func (c Coordinate) GetNeighbors() []Coordinate {
	neighbors := make([]Coordinate, 6)
	for i := range neighbors {
		current := c.Copy()
		current.Move(movesToNeighbors[i])
		neighbors[i] = current
	}
	return neighbors
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c[0], c[1])
}
