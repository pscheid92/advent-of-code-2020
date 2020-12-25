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

type Storage map[Coordinate]Tile

func (s Storage) FlipTile(coordinate Coordinate) {
	if tile, ok := s[coordinate]; ok {
		s[coordinate] = tile.Flip()
	} else {
		s[coordinate] = BLACK
	}
}

func (s Storage) CountTiles(side Tile) int {
	sum := 0
	for _, tile := range s {
		if tile == side {
			sum += 1
		}
	}
	return sum
}
