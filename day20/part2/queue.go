package main

type TileQueue []Tile

func (tq *TileQueue) Push(tile Tile) {
	// do nothing if already present
	for _, x := range *tq {
		if x.ID == tile.ID {
			return
		}
	}

	// else append
	*tq = append(*tq, tile)
}

func (tq *TileQueue) Pop() Tile {
	tile := (*tq)[0]
	*tq = (*tq)[1:]
	return tile
}

func (tq TileQueue) NotEmpty() bool {
	return len(tq) > 0
}
