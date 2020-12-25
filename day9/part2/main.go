package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day9/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	numbers, err := helpers.ConvertLinesToNumbers(lines)
	if err != nil {
		log.Fatalln(err)
	}

	target := 375054920
	group := FindConsecutiveGroupByPrefixSum(target, numbers)
	min, max := helpers.FindMinAndMax(group)
	fmt.Printf("solution code: %d\n", min+max)
}

func FindConsecutiveGroupByBruteForce(target int, numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += numbers[k]
			}

			if sum == target {
				return numbers[i : j+1]
			}
		}
	}
	return nil
}

func FindConsecutiveGroupByPrefixSum(target int, numbers []int) []int {
	n := len(numbers)

	// create prefix sums
	prefixSums := make([]int, n)
	prefixSums[0] = numbers[0]
	for i := 1; i < n; i++ {
		prefixSums[i] = prefixSums[i-1] + numbers[i]
	}

	// find group
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prefixSums[j]-prefixSums[i] == target {
				return numbers[i+1 : j+1]
			}
		}
	}

	return nil
}

func IsAtIndexValid(index int, preambleSize int, list []int) bool {
	if index < preambleSize {
		return false
	}

	current := list[index]
	candidates := list[index-preambleSize : index]

	for i, x := range candidates {
		for j, y := range candidates {
			// not allowed to use the same number twice
			if i == j {
				continue
			}

			// matching sum found? number is valid!
			if x+y == current {
				return true
			}
		}
	}

	// we found no matching sum :-(
	return false
}
