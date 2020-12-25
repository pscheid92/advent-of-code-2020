package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoordinate_Move(t *testing.T) {
	// move to west
	c := NewCoordinate()
	c.Move(W)
	assert.Equal(t, Coordinate{-1, 0}, c)

	// run circle in east direction
	c = NewCoordinate()
	c.Move(NE)
	c.Move(SE)
	c.Move(W)
	assert.Equal(t, Coordinate{0, 0}, c)

	// run circle in west direction
	c = NewCoordinate()
	c.Move(NW)
	c.Move(SW)
	c.Move(E)
	assert.Equal(t, Coordinate{0, 0}, c)
}

func TestCoordinate_MovePath(t *testing.T) {
	// move to west
	c := NewCoordinate()
	c.MovePath([]Move{W, W, W})
	assert.Equal(t, Coordinate{-3, 0}, c)

	// run circle in east direction
	c = NewCoordinate()
	c.MovePath([]Move{NE, SE, W})
	assert.Equal(t, Coordinate{0, 0}, c)

	// run circle in west direction
	c = NewCoordinate()
	c.MovePath([]Move{NW, SW, E})
	assert.Equal(t, Coordinate{0, 0}, c)
}
