package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day6/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)

	sum := 0
	for _, g := range groups {
		sum += CountCommonAnswers(g)
	}
	fmt.Printf("solution code is %d\n", sum)
}

// CountCommonAnswers counts common answers among all people within a group.
//
// Within a group the answers of each person is in a separate line.
// We count the occurrences of every answer (single character) in every line.
//
// Example:
//		abc
//		a
//
// Result:
//		a: 2, b: 1, c: 1
//
// This group consists of two persons (two lines).
// Thus, if any answer was given two times, it is a common answer among all (2) persons.
// And therefore is counted as a common answer.
//
func CountCommonAnswers(group []string) int {
	// count occurrences of answers among people in group
	letterCounter := make(map[rune]int)
	for _, line := range group {
		for _, c := range line {
			// ignore invalid symbols
			if c < 'a' || 'z' < c {
				continue
			}

			letterCounter[c]++
		}
	}

	// determine and count answers given by all people in group
	count := 0
	for _, n := range letterCounter {
		if n == len(group) {
			count++
		}
	}
	return count
}
