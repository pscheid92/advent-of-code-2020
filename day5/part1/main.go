package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	maxID := 0
	for _, l := range lines {
		id := DecodeSeat(l)
		if id > maxID {
			maxID = id
		}
	}

	fmt.Printf("solution code: %d\n", maxID)
}

func DecodeSeat(line string) int {
	row := translateToBinary(line[:7])
	col := translateToBinary(line[7:])
	id := row*8 + col
	return id
}

func translateToBinary(word string) int {
	i := 0
	for _, c := range word {
		i <<= 1
		if c == 'R' || c == 'B' {
			i |= 0b00000001
		}
	}
	return i
}
