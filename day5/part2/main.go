package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"sort"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// generate ids and sort them
	ids := make([]int, len(lines))
	for i, l := range lines {
		ids[i] = DecodeSeat(l)
	}
	sort.Ints(ids)

	// check if id+1 is in list (if not we found the gap)
	for i, id := range ids {
		if id+1 != ids[i+1] {
			fmt.Printf("solution code: %d\n", id+1)
			return
		}
	}
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
