package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"sort"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day10/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	jolts, err := helpers.ConvertLinesToNumbers(lines)
	if err != nil {
		log.Fatalln(err)
	}

	combinations := CalculateCombinations(jolts)
	fmt.Printf("solution code: %d\n", combinations)
}

func CalculateCombinations(originalJolts []int) int {
	// create a copy to prevent side-effects for caller
	jolts := make([]int, len(originalJolts))
	copy(jolts, originalJolts)

	// sort the jolts and append 0 as start and max+3 as ending jolt
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	jolts = append([]int{0}, jolts...)

	// prepare memoization memory
	memory := make(map[int]int)
	n := len(jolts)

	// prepare recursive function
	var recursiveFunction func(int) int
	recursiveFunction = func(current int) int {
		if v, ok := memory[current]; ok {
			return v
		}

		if current >= n-1 {
			return 1
		}

		possibilities := 0
		for i := current + 1; i < n && jolts[i] <= jolts[current]+3; i++ {
			possibilities += recursiveFunction(i)
		}

		memory[current] = possibilities
		return possibilities
	}

	// start actual calculation
	return recursiveFunction(0)
}
