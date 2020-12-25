package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTileQueue_Push(t *testing.T) {
	tileA := NewTile(1, Original, RotatedDown, []string{"x"})
	tileB := NewTile(2, Original, RotatedDown, []string{"x"})
	tileC := NewTile(2, Original, RotatedDown, []string{"x"})

	q := TileQueue{}
	assert.Len(t, q, 0)

	q.Push(tileA)
	q.Push(tileB)
	q.Push(tileC) // will not work because id=2 already present
	assert.Len(t, q, 2)
}

func TestTileQueue_Pop(t *testing.T) {
	tileA := NewTile(1, Original, RotatedDown, []string{"x"})
	tileB := NewTile(2, Original, RotatedDown, []string{"x"})

	q := TileQueue{}
	assert.Panics(t, func() { q.Pop() })

	q.Push(tileA)
	q.Push(tileB)

	assert.Equal(t, tileA.ID, q.Pop().ID)
	assert.Equal(t, tileB.ID, q.Pop().ID)
}
