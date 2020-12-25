package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	forest, err := helpers.ReadLineByLineFromFile("day3/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	count := WalkThroughForest(3, 1, forest)
	fmt.Printf("solution is %d\n", count)
}

const TreeMarker = '#'

func WalkThroughForest(slopeX int, slopeY int, forest []string) int {
	height := len(forest)
	width := len(forest[0])

	count := 0
	x := slopeX
	y := slopeY

	for y < height {
		if forest[y][x] == TreeMarker {
			count++
		}

		x = (x + slopeX) % width
		y += slopeY
	}

	return count
}
