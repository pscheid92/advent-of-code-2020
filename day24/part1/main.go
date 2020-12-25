package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day24/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	paths, err := ParsePaths(lines)
	if err != nil {
		log.Fatalln(err)
	}

	storage := make(Storage)
	for _, p := range paths {
		c := NewCoordinate()
		c.MovePath(p)
		storage.FlipTile(c)
	}

	blackTiles := storage.CountTiles(BLACK)
	fmt.Printf("solution code: %d\n", blackTiles)
}
