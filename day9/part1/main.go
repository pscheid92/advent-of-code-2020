package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day9/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	numbers, err := helpers.ConvertLinesToNumbers(lines)
	if err != nil {
		log.Fatalln(err)
	}

	preambleSize := 25
	for i := preambleSize; i < len(numbers); i++ {
		if !IsAtIndexValid(i, preambleSize, numbers) {
			fmt.Printf("found invalid number '%d' at index %d\n", numbers[i], i)
			fmt.Printf("solution code is: %d\n", numbers[i])
			return
		}
	}

	log.Fatalln("no solution found")
}

func IsAtIndexValid(index int, preambleSize int, list []int) bool {
	if index < preambleSize {
		return false
	}

	current := list[index]
	candidates := list[index-preambleSize : index]

	for i, x := range candidates {
		for j, y := range candidates {
			// not allowed to use the same number twice
			if i == j {
				continue
			}

			// matching sum found? number is valid!
			if x+y == current {
				return true
			}
		}
	}

	// we found no matching sum :-(
	return false
}
