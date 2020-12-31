package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"os"
	"sort"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day1/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	entries, err := helpers.ConvertLinesToNumbers(lines)
	if err != nil {
		log.Fatalln(err)
	}

	results, ok := FindEntriesSummingTo(entries, 2020)
	if !ok {
		fmt.Println("no matching entries found")
		os.Exit(1)
	}

	fmt.Printf("matching entries are %d and %d\n", results[0], results[1])
	fmt.Printf("answer code is %d\n", results[0]*results[1])
}

func FindEntriesSummingTo(entries []int, target int) ([2]int, bool) {
	sort.Ints(entries)

	i := 0
	j := len(entries) - 1

	for i < j && entries[i]+entries[j] != target {
		// to big? narrow from right
		if entries[i]+entries[j] > target {
			j--
		}

		// to small? narrow from left
		if entries[i]+entries[j] < target {
			i++
		}
	}

	return [2]int{entries[i], entries[j]}, i < j
}
