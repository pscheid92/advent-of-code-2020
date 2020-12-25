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

	results, ok := FindTreeEntriesSummingTo(entries, 2020)
	if !ok {
		fmt.Println("no matching entries found")
		os.Exit(1)
	}

	fmt.Printf("matching entries are %d, %d and %d\n", results[0], results[1], results[2])
	fmt.Printf("answer code is %d\n", results[0]*results[1]*results[2])
}

func FindTwoEntriesSummingTo(entries []int, target int) ([2]int, bool) {
	// overall complexity O(n) with n = len(entries)

	i := 0
	j := len(entries) - 1

	// ignore all, that are too big alone
	for j >= 0 && entries[j] > target {
		j--
	}

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

	if j <= i {
		return [2]int{}, false
	}

	return [2]int{entries[i], entries[j]}, true
}

func FindTreeEntriesSummingTo(entries []int, target int) ([3]int, bool) {
	// overall complexity O(n^2)

	// O(n log n)
	sort.Sort(sort.IntSlice(entries))

	// O(n) for loop and O(n) for work
	for j := len(entries) - 1; j > 0; j-- {
		slice := entries[:j]
		results, ok := FindTwoEntriesSummingTo(slice, target-entries[j])
		if ok {
			return [3]int{results[0], results[1], entries[j]}, true
		}
	}

	// failed to find entries
	return [3]int{0, 0, 0}, false
}
