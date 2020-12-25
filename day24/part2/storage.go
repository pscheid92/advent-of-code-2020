package main

type Storage map[Coordinate]Tile

func (s Storage) Get(coordinate Coordinate) Tile {
	if tile, ok := s[coordinate]; ok {
		return tile
	} else {
		s[coordinate] = WHITE
		return WHITE
	}
}

func (s Storage) Set(coordinate Coordinate, tile Tile) {
	s[coordinate] = tile
}

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

func (s Storage) GetInterestingCoordinates() []Coordinate {
	set := NewCoordinateSet(len(s) * 6)
	for c := range s {
		set.Add(c, c.GetNeighbors()...)
	}
	return set.ToSlice()
}

func (s Storage) CountTilesInNeighborhood(coordinate Coordinate, tile Tile) int {
	sum := 0
	for _, neighbor := range coordinate.GetNeighbors() {
		if s.Get(neighbor) == tile {
			sum += 1
		}
	}
	return sum
}

func (s Storage) InitAccordingToPaths(paths [][]Move) {
	for _, p := range paths {
		c := NewCoordinate()
		c.MovePath(p)
		s.FlipTile(c)
	}
}

func (s Storage) ApplyRulesNTimes(n int) {
	type changeInstruction struct {
		coordinate Coordinate
		tile       Tile
	}

	for i := 0; i < n; i++ {
		interestingCoordinates := s.GetInterestingCoordinates()
		instructions := make([]changeInstruction, 0)

		// phase 1: collect changes
		for _, c := range interestingCoordinates {
			blacks := s.CountTilesInNeighborhood(c, BLACK)
			currentTile := s.Get(c)

			if currentTile == BLACK && (blacks == 0 || blacks > 2) {
				instructions = append(instructions, changeInstruction{c, WHITE})
			}

			if currentTile == WHITE && blacks == 2 {
				instructions = append(instructions, changeInstruction{c, BLACK})
			}
		}

		// phase 2: apply changes
		for _, instruction := range instructions {
			s.Set(instruction.coordinate, instruction.tile)
		}
	}
}
