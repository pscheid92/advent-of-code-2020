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

	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	encounters := WalkThroughMultipleForests(slopes, forest)
	solution := helpers.Multiply(encounters)
	fmt.Printf("solution is %d\n", solution)
}

const TreeMarker = '#'

func WalkThroughMultipleForests(slopes [][2]int, forest []string) []int {
	encounters := make([]int, len(slopes))
	for i, slope := range slopes {
		encounters[i] = WalkThroughForest(slope[0], slope[1], forest)
	}
	return encounters
}

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
