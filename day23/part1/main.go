package main

import (
	"container/ring"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "963275481"
	moves := 100

	// create game and play
	currentCup, min, max := CreateGame(input)
	currentCup = PlayNMoves(moves, currentCup, min, max)

	// generate output aka solution code
	output, err := GenerateOutput(currentCup)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("solution code: %s\n", output)
}

func CreateGame(input string) (*ring.Ring, int, int) {
	n := len(input)
	min := math.MaxInt32
	max := math.MinInt32

	// convert symbols to numbers
	// on the fly, find minimum and maximum value
	numbers := make([]int, n)
	for i, x := range input {
		numbers[i] = int(x - '0')

		if numbers[i] < min {
			min = numbers[i]
		}

		if numbers[i] > max {
			max = numbers[i]
		}
	}

	// create ring and fill it with the numbers
	cup := ring.New(n)
	for _, x := range numbers {
		cup.Value = x
		cup = cup.Next()
	}

	return cup, min, max
}

func PlayMove(currentCup *ring.Ring, min int, max int) *ring.Ring {
	// pick three cups after current cup
	pickedUp := currentCup.Unlink(3)

	// search destination cup and insert picked up cups after it
	destinationCup := SearchDestinationCup(currentCup, min, max)
	destinationCup.Link(pickedUp)

	// return next current cup
	return currentCup.Next()
}

func PlayNMoves(n int, currentCup *ring.Ring, min int, max int) *ring.Ring {
	for i := 0; i < n; i++ {
		currentCup = PlayMove(currentCup, min, max)
	}
	return currentCup
}

func SearchDestinationCup(start *ring.Ring, min, max int) *ring.Ring {
	searchValue := start.Value.(int)
	for {
		// calculate next target value
		searchValue = searchValue - 1
		if searchValue < min {
			searchValue = max
		}

		// search for target value
		if cup, found := FindCupWithValue(start, searchValue); found {
			return cup
		}
	}
}

func GenerateOutput(startCup *ring.Ring) (string, error) {
	cup, found := FindCupWithValue(startCup, 1)
	if !found {
		return "", fmt.Errorf("no cup with label 1 found")
	}

	var builder strings.Builder
	for current := cup.Next(); current != cup; current = current.Next() {
		value := current.Value.(int)
		builder.WriteString(strconv.Itoa(value))
	}
	return builder.String(), nil
}

func FindCupWithValue(startCup *ring.Ring, searchedValue int) (*ring.Ring, bool) {
	// check if start cup contains the searched value
	// do it here, because the later loop skips the start cup
	if startCup.Value.(int) == searchedValue {
		return startCup, true
	}

	// search remaining cups for searched value
	for current := startCup.Next(); current != startCup; current = current.Next() {
		if current.Value.(int) == searchedValue {
			return current, true
		}
	}

	return nil, false
}
