package main

import (
	"fmt"
)

type Coordinate [2]int

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

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c[0], c[1])
}
