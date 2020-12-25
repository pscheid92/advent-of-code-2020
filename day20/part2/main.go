package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	tiles, err := SolvePartOne("day20/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	image := RebuildImage(tiles)
	variations := Variations([]Tile{image})

	correctImage := image
	monsterCounter := 0

	for _, x := range variations {
		monsterCounter = CountMonsters(x)
		if monsterCounter != 0 {
			correctImage = x
			break
		}
	}

	waveCounter := 0
	for _, row := range correctImage.Data {
		for _, c := range row {
			if c == '#' {
				waveCounter++
			}
		}
	}

	result := waveCounter - 15*monsterCounter
	fmt.Printf("solution code: %d\n", result)
}

func SolvePartOne(path string) ([]Tile, error) {
	lines, err := helpers.ReadLineByLineFromFile(path)
	if err != nil {
		return nil, err
	}
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)
	tiles, err := ReadTiles(groups)
	if err != nil {
		return nil, err
	}
	return ConnectTiles(tiles), nil
}

func Variations(tiles []Tile) []Tile {
	variations := make([]Tile, 0, 8*len(tiles))

	// just unrolled the possible variations (instead of using loops)
	for _, t := range tiles {
		tmp := make([]Tile, 8)
		tmp[0] = t
		tmp[1] = t.Rotate()
		tmp[2] = tmp[1].Rotate()
		tmp[3] = tmp[2].Rotate()
		tmp[4] = t.Flip()
		tmp[5] = tmp[4].Rotate()
		tmp[6] = tmp[5].Rotate()
		tmp[7] = tmp[6].Rotate()
		variations = append(variations, tmp...)
	}

	return variations
}

func ConnectTiles(originalTiles []Tile) []Tile {
	result := make([]Tile, 0, len(originalTiles))
	remainingTiles := Variations(originalTiles[1:])
	queue := TileQueue{originalTiles[0]}

	for queue.NotEmpty() {
		current := queue.Pop()

		// check against all remaining tiles (or until all neighbors for current are found)
		for i := 0; i < len(remainingTiles) && current.CountNeighbors() < 4; i++ {
			other := remainingTiles[i]

			// if tiles do not match, try next one
			if matched := MatchTiles(&current, &other); !matched {
				continue
			}

			// add to queue if not already present
			queue.Push(other)

			// remove variants of remaining
			// DANGER: implicitly assumes that all further occurences are AFTER this
			newRemainingTiles := make([]Tile, 0)
			for _, x := range remainingTiles {
				if other.ID != x.ID {
					newRemainingTiles = append(newRemainingTiles, x)
				}
			}
			remainingTiles = newRemainingTiles
		}

		result = append(result, current)
	}

	// run over remaining tiles and build missing connections
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			_ = MatchTiles(&result[i], &result[j])
		}
	}
	return result
}

func RebuildImage(tiles []Tile) Tile {
	// remove boarder of all tiles
	for i := range tiles {
		tiles[i].RemoveBorder()
	}

	// create a lookup table
	lookup := make(map[int]Tile, len(tiles))
	for _, x := range tiles {
		lookup[x.ID] = x
	}

	// find corner piece (upper left corner)
	var corner Tile
	for _, tile := range tiles {
		if tile.neighbors[UpperSide] == 0 && tile.neighbors[LeftSide] == 0 {
			corner = tile
			break
		}
	}

	// crawl from corner down to get left border tiles
	leftBorder := FindTilesInDirection(corner, LowerSide, lookup)

	// preallocate data
	n := len(corner.Data)
	merged := make([]string, n*len(leftBorder))

	// merge image:
	// run through left border tiles
	// for each tile search for the tiles building this row
	// merge them together in a new string slice
	for i, left := range leftBorder {
		for _, tile := range FindTilesInDirection(left, RightSide, lookup) {
			for j, row := range tile.Data {
				merged[i*n+j] = merged[i*n+j] + row
			}
		}
	}

	// put merged data bag into one big monster tile
	return NewTile(0, Original, RotatedUp, merged)
}

func FindTilesInDirection(corner Tile, side Side, lookup map[int]Tile) []Tile {
	result := make([]Tile, 0)
	current := corner
	for {
		result = append(result, current)
		if next, found := lookup[current.GetNeighbor(side)]; !found {
			break
		} else {
			current = next
		}
	}
	return result
}

func CountMonsters(image Tile) int {
	stencil := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	length := len(stencil)
	width := len(stencil[1])

	counter := 0
	for i := 0; i < len(image.Data)-length; i++ {
		for j := 0; j < len(image.Data[0])-width; j++ {

			allMatch := true
			for k := 0; allMatch && k < length; k++ {
				for l := 0; l < width; l++ {
					if stencil[k][l] == '#' && image.Data[i+k][j+l] != '#' {
						allMatch = false
						break
					}
				}
			}

			if allMatch {
				counter++
			}
		}
	}

	return counter
}
