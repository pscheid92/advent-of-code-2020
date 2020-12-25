package main

import "fmt"

func main() {
	startSequence := []int{9, 12, 1, 4, 17, 0, 18}
	result := PlayGame(startSequence, 30000000)
	fmt.Printf("solution code: %d\n", result)
}

func PlayGame(startSequence []int, finalRound int) int {
	// storage
	storage := make(map[int]int, finalRound/2)
	n := len(startSequence) - 1

	for i := 0; i < n; i++ {
		current := startSequence[i]
		storage[current] = i
	}

	// state
	var result int
	var next int
	current := startSequence[n]

	// play rounds
	for round := n; round < finalRound; round++ {
		if lastSeen, found := storage[current]; !found {
			next = 0
		} else {
			next = round - lastSeen
		}

		storage[current] = round
		result = current
		current = next
	}

	return result
}
