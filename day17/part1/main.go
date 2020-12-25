package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day17/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	count := RunSimulation(lines, 6)
	fmt.Printf("solution code: %d\n", count)
}

func RunSimulation(input []string, steps int) int {
	world := InitWorld(input)
	world.SimulateSteps(steps)
	return world.CountActive()
}

