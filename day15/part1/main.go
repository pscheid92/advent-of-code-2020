package main

import "fmt"

func main() {
	startSequence := []int{9, 12, 1, 4, 17, 0, 18}
	result := PlayGame(startSequence, 2020)
	fmt.Printf("solution code: %d\n", result)
}

func PlayGame(startSequence []int, finalRound int) int {
	// copy
	numbers := make([]int, len(startSequence))
	copy(numbers, startSequence)

	// init
	for round := len(startSequence) - 1; round < finalRound; round++ {
		current := numbers[round]

		found := -1
		for i := round - 1; i >= 0; i-- {
			if numbers[i] == current {
				found = i
				break
			}
		}

		if found == -1 {
			numbers = append(numbers, 0)
		} else {
			numbers = append(numbers, round-found)
		}
	}

	return numbers[finalRound-1]
}
