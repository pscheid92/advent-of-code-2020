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

	diffs := CountJoltDifferences(jolts)
	fmt.Printf("solution code: %d\n", diffs[1]*diffs[3])
}

func CountJoltDifferences(jolts []int) map[int]int {
	differences := make(map[int]int)
	n := len(jolts)
	sort.Ints(jolts)

	outlet := []int{0}
	device := jolts[n-1] + 3
	jolts = append(outlet, jolts...)
	jolts = append(jolts, device)

	for i := 1; i < n+2; i++ {
		diff := jolts[i] - jolts[i-1]
		if _, ok := differences[diff]; ok {
			differences[diff]++
		} else {
			differences[diff] = 1
		}
	}

	return differences
}
