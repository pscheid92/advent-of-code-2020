package main

type Tile int

const (
	WHITE Tile = iota
	BLACK
)

func (s Tile) Flip() Tile {
	return (s + 1) & 1
}

func (s Tile) String() string {
	switch s {
	case WHITE:
		return "WHITE"
	case BLACK:
		return "BLACK"
	default:
		panic("kaboom")
	}
}
