package main

import (
	"container/ring"
	"fmt"
	"log"
	"math"
)

func main() {
	startSequence := "963275481"
	size := 1_000_000
	moves := 10_000_000

	currentCup, min, max, lookup := CreateGame(startSequence, size)
	_ = PlayNMoves(moves, currentCup, min, max, lookup)

	// generate output aka solution code
	output, err := CalculateSolutionCode(lookup)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("solution code: %d\n", output)
}

func CreateGame(startSequence string, size int) (*ring.Ring, int, int, []*ring.Ring) {
	cups := ring.New(size)
	lookup := make([]*ring.Ring, size+1)

	// convert start sequence to numbers
	// find minimum and maximum value on the fly
	numbers, n, min, max := convertSequence(startSequence)

	// write start sequence to cups
	for _, x := range numbers {
		cups.Value = x
		lookup[x] = cups
		cups = cups.Next()
	}

	// fill remaining values of ring starting with max+1
	// at the end, we have filled all cups [start sequence..., max+1, max+1, ...]
	nextValue := max
	for i := n; i < size; i++ {
		nextValue++
		cups.Value = nextValue
		lookup[nextValue] = cups
		cups = cups.Next()
	}

	// underway nextValue always contains the maximum value of all cups
	return cups, min, nextValue, lookup
}

func PlayNMoves(n int, currentCup *ring.Ring, min int, max int, lookup []*ring.Ring) *ring.Ring {
	for i := 0; i < n; i++ {
		currentCup = PlayMove(currentCup, min, max, lookup)
	}
	return currentCup
}

func PlayMove(currentCup *ring.Ring, min int, max int, lookup []*ring.Ring) *ring.Ring {
	// pick three cups after current cup
	pickedUp := currentCup.Unlink(3)
	forbidden := [3]int{
		pickedUp.Value.(int),
		pickedUp.Next().Value.(int),
		pickedUp.Next().Next().Value.(int),
	}

	// search destination cup and insert picked up cups after it
	destinationCup := SearchDestinationCup(currentCup, min, max, lookup, forbidden)
	destinationCup.Link(pickedUp)

	// return next current cup
	return currentCup.Next()
}

func SearchDestinationCup(start *ring.Ring, min, max int, lookup []*ring.Ring, forbidden [3]int) *ring.Ring {
	searchValue := start.Value.(int)
	for {
		// calculate next target value
		searchValue = searchValue - 1
		if searchValue < min {
			searchValue = max
		}

		// if forbidden value ... try next
		if searchValue == forbidden[0] || searchValue == forbidden[1] || searchValue == forbidden[2] {
			continue
		}

		// actually try to lookup cup
		cup := lookup[searchValue]
		if cup != nil {
			return cup
		}
	}
}

func CalculateSolutionCode(lookup []*ring.Ring) (int, error) {
	cup := lookup[1]
	if cup == nil {
		return 0, fmt.Errorf("no cup with label 1 found")
	}

	first := cup.Next().Value.(int)
	second := cup.Next().Next().Value.(int)

	return first * second, nil
}

func convertSequence(sequence string) ([]int, int, int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	result := make([]int, len(sequence))
	for i, x := range sequence {
		result[i] = int(x - '0')

		if result[i] < min {
			min = result[i]
		}

		if result[i] > max {
			max = result[i]
		}
	}

	return result, len(sequence), min, max
}
