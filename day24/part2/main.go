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
	storage.InitAccordingToPaths(paths)

	storage.ApplyRulesNTimes(100)
	sum := storage.CountTiles(BLACK)

	fmt.Printf("solution code: %d\n", sum)
}
