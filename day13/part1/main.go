package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day13/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	now, busses, err := parseInputLines(lines)
	if err != nil {
		log.Fatalln(err)
	}

	arrivalTime, nextBus := CalculateNextBus(now, busses)
	fmt.Printf("solution code: %d\n", arrivalTime*nextBus)
}

func parseInputLines(lines []string) (int, []int, error) {
	now, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}

	busses := make([]int, 0)
	for _, b := range strings.Split(lines[1], ",") {
		if b == "x" {
			continue
		}

		id, err := strconv.Atoi(b)
		if err != nil {
			return 0, nil, err
		}

		busses = append(busses, id)
	}

	return now, busses, nil
}

func CalculateNextBus(now int, busses []int) (int, int) {
	minimumTime := math.MaxInt32
	minimumBus := math.MaxInt32

	for _, bus := range busses {
		time := bus - (now % bus)
		if time < minimumTime {
			minimumTime = time
			minimumBus = bus
		}
	}

	return minimumTime, minimumBus
}
