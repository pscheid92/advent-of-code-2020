package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day16/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	specifications, myTicket, nearbyTickets, err := ExtractInput(lines)
	if err != nil {
		log.Fatalln(err)
	}
	nearbyTickets = RemoveInvalidTickets(nearbyTickets, specifications)

	m := InitMatrix(specifications, nearbyTickets)
	orderedFields, err := SolveFieldOrder(m, 1_000)
	if err != nil {
		log.Fatalln(err)
	}

	mul := 1
	for i, field := range orderedFields {
		if strings.HasPrefix(field, "departure") {
			mul *= myTicket[i]
		}
	}

	fmt.Printf("solution code: %d\n", mul)
}

func SolveFieldOrder(matrix Matrix, maxIter uint) ([]string, error) {
	iter := uint(0)
	orderedFields := make([]string, len(matrix.Columns))

	for matrix.N != 0 || matrix.M != 0 {
		// guard against endless loop
		if iter >= maxIter {
			return nil, fmt.Errorf("reached maximum allowed iterations")
		}
		iter++

		// check from row view
		for i := 0; i < matrix.N; i++ {
			sum, last := matrix.RowSum(i)
			if sum == 1 {
				// save order
				spec := matrix.Specifications[i]
				col := matrix.Columns[last]
				orderedFields[col] = spec.Name

				matrix.RemoveRow(i)
				matrix.RemoveCol(last)
				continue
			}
		}

		// check from col view
		for j := 0; j < matrix.M; j++ {
			sum, last := matrix.ColSum(j)
			if sum == 1 {
				// save order
				spec := matrix.Specifications[last]
				col := matrix.Columns[j]
				orderedFields[col] = spec.Name

				matrix.RemoveRow(last)
				matrix.RemoveCol(j)
				continue
			}
		}
	}

	return orderedFields, nil
}
