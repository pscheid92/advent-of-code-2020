package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"regexp"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day7/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rules, ok := BuildRuleTable(lines)
	if !ok {
		log.Fatalln("error building rules table")
	}

	outerBags := DeterminePossibleOuterBags("shiny gold", rules)
	fmt.Printf("solution code: %d\n", len(outerBags))
}

var (
	OuterRulePattern = regexp.MustCompile(`^([a-z ]+) bags contain (.*)\.$`)
	InnerRulePattern = regexp.MustCompile(`^\d+ ([a-z ]+) bags?$`)
)

func DeterminePossibleOuterBags(targetBag string, rules map[string][]string) []string {
	queue := []string{targetBag}
	for i := 0; i < len(queue); i++ {
		for bag, allowed := range rules {
			// a map to memorise already seen bags could save some cpu cycles (trade-off memory)
			if !helpers.StringSliceContains(bag, queue) && helpers.StringSliceContains(queue[i], allowed) {
				queue = append(queue, bag)
			}
		}
	}
	return queue[1:]
}

func BuildRuleTable(lines []string) (map[string][]string, bool) {
	table := make(map[string][]string)
	for _, l := range lines {
		bag, allowed, ok := ExtractRule(l)
		if !ok {
			return nil, false
		}

		if _, ok := table[bag]; ok {
			return nil, false
		}

		table[bag] = allowed
	}

	return table, true
}

func ExtractRule(line string) (string, []string, bool) {
	// break rule into bag and inner part
	match := OuterRulePattern.FindStringSubmatch(line)
	if match == nil {
		return "", nil, false
	}

	bag := match[1]
	innerText := match[2]
	allowed := make([]string, 0)

	if innerText == "no other bags" {
		return bag, allowed, true
	}

	parts := strings.Split(innerText, ", ")
	for _, p := range parts {
		match = InnerRulePattern.FindStringSubmatch(p)
		if match == nil {
			return "", nil, false
		}
		allowed = append(allowed, match[1])
	}

	return bag, allowed, true
}
