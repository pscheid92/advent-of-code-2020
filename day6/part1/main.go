package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	// read input
	lines, err := helpers.ReadLineByLineFromFile("day6/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	groups := helpers.StackMultilineSeparatedByEmptyOne(lines)

	sum := 0
	for _, g := range groups {
		sum += CountUniqueLetters(g)
	}
	fmt.Printf("solution code is %d\n", sum)
}

func CountUniqueLetters(group string) int {
	counter := make(map[rune]byte)

	for _, c := range group {
		// ignore invalid symbols
		if c < 'a' || 'z' < c {
			continue
		}

		if _, ok := counter[c]; ok {
			continue
		} else {
			counter[c] = 0
		}
	}

	return len(counter)
}
