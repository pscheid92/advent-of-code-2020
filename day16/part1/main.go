package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day16/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	specifications, _, nearbyTickets, err := ExtractInput(lines)
	if err != nil {
		log.Fatalln(err)
	}

	invalidColumns := make([]int, 0)
	for _, ticket := range nearbyTickets {
		isValid, invalidColumn := CheckTicket(ticket, specifications)
		if !isValid {
			invalidColumns = append(invalidColumns, invalidColumn)
		}
	}

	errorRate := helpers.Sum(invalidColumns)
	fmt.Printf("solution code: %d\n", errorRate)
}

func CheckTicket(ticket Ticket, specifications []FieldSpecification) (bool, int) {
	for _, column := range ticket {
		// if any spec matches, this column is not totally invalid
		anySpecMatches := false
		for _, s := range specifications {
			if s.IsValid(column) {
				anySpecMatches = true
				break
			}
		}

		// a column without any matching spec means this ticket is totally invalid
		// we can abort the check and remember the value that made it invalid
		if !anySpecMatches {
			return false, column
		}
	}
	return true, -1
}

type Ticket []int

type FieldSpecification struct {
	Name   string
	ranges [2][2]int
}

func (fs FieldSpecification) IsValid(x int) bool {
	first := fs.ranges[0][0] <= x && x <= fs.ranges[0][1]
	second := fs.ranges[1][0] <= x && x <= fs.ranges[1][1]
	return first || second
}
