package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day20/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)
	tiles, err := ReadTiles(groups)
	if err != nil {
		log.Fatalln(err)
	}

	result := ConnectTiles(tiles)
	corners := FindCorners(result)
	if len(corners) != 4 {
		log.Fatalf("found %d corner tiles, expected 4", len(corners))
	}

	mul := 1
	for _, c := range corners {
		mul *= c.ID
	}

	fmt.Printf("soultion code: %d\n", mul)
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

func FindCorners(tiles []Tile) []Tile {
	result := make([]Tile, 0, 4)
	for _, t := range tiles {
		if t.CountNeighbors() == 2 {
			result = append(result, t)
		}
	}
	return result
}
