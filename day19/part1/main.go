package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day19/part1/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)
	lookup := ParseRules(groups[0])

	var count int
	for _, word := range groups[1] {
		valid := IsValidWord(word, lookup, "0")
		if valid {
			count++
		}
	}

	fmt.Printf("solution code: %d\n", count)
}

func ParseRules(lines []string) map[string]StringSet {
	// preallocate with len(lines) as a rule of thumb
	lookup := make(map[string]StringSet, len(lines))

	for _, l := range lines {
		outerParts := strings.Split(l, ":")
		ruleKey := outerParts[0]

		innerParts := strings.Split(outerParts[1], "|")
		for _, p := range innerParts {
			p = strings.TrimSpace(p)
			p = strings.ReplaceAll(p, `"`, "")
			p = strings.ReplaceAll(p, " ", ",")

			if _, found := lookup[p]; !found {
				lookup[p] = NewStringSet()
			}
			lookup[p].Add(ruleKey)
		}
	}

	return lookup
}

func IsValidWord(word string, lookup map[string]StringSet, startRule string) bool {
	n := len(word)

	// init table (n x n)
	table := make([][]StringSet, n)
	for i := 0; i < n; i++ {
		table[i] = make([]StringSet, n)
		for j := 0; j < n; j++ {
			table[i][j] = NewStringSet()
		}
	}

	// write terminal rules into first line
	for i, letter := range word {
		rules, found := lookup[string(letter)]
		if !found {
			return false
		}
		table[0][i] = rules
	}

	// start filling table
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {

			// check previously filled table according to cyk pattern
			for k := 0; k < i; k++ {
				// look into table fields
				x := table[k][j]
				y := table[i-k-1][j+k+1]

				// build combinations of table fields
				// check if there is a rule for each combination
				// if so, we save it for later runs
				for _, p := range BuildCrossProduct(x, y) {
					if rules, found := lookup[p]; found {
						table[i][j].Union(rules)
					}
				}
			}
		}
	}

	// word is valid if last field in triangle contains start rule
	rules := table[n-1][0]
	return rules.Contains(startRule)
}

func BuildCrossProduct(xs StringSet, ys StringSet) []string {
	// no xs? return plain ys
	if xs.Len() == 0 {
		return ys.GetAll()
	}

	// no ys? return plain xs
	if ys.Len() == 0 {
		return xs.GetAll()
	}

	// both present? build cross product
	result := make([]string, 0, xs.Len()*ys.Len())
	for _, x := range xs.GetAll() {
		for _, y := range ys.GetAll() {
			result = append(result, x+","+y)
		}
	}
	return result
}
