package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"regexp"
	"strconv"
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

	count := CalculateContainedBags("shiny gold", rules)
	fmt.Printf("solution code: %d\n", count)
}

var (
	OuterRulePattern = regexp.MustCompile(`^([a-z ]+) bags contain (.*)\.$`)
	InnerRulePattern = regexp.MustCompile(`^(\d+) ([a-z ]+) bags?$`)
)

type Rule struct {
	bag        string
	allowed    []string
	quantities []int
}

func CalculateContainedBags(targetBag string, rules map[string]Rule) int {
	rule := rules[targetBag]

	// this targetBag contains no other bags
	if len(rule.allowed) == 0 {
		return 0
	}

	possibleBags := 0
	for i := 0; i < len(rule.allowed); i++ {
		currentSubBag := rule.allowed[i]
		quantityOfCurrentSubBag := rule.quantities[i]

		// count the subBags themselfs, since they are in the targetBag
		possibleBags += quantityOfCurrentSubBag

		// additionally add the bags within the subBags
		bagsWithinASingleSubBag := CalculateContainedBags(currentSubBag, rules)
		possibleBags += quantityOfCurrentSubBag * bagsWithinASingleSubBag
	}

	return possibleBags
}

func BuildRuleTable(lines []string) (map[string]Rule, bool) {
	table := make(map[string]Rule)
	for _, l := range lines {
		rule, ok := ExtractRule(l)
		if !ok {
			return nil, false
		}

		if _, ok := table[rule.bag]; ok {
			return nil, false
		}

		table[rule.bag] = rule
	}

	return table, true
}

func ExtractRule(line string) (Rule, bool) {
	// break rule into bag and inner part
	match := OuterRulePattern.FindStringSubmatch(line)
	if match == nil {
		return Rule{}, false
	}

	rule := Rule{
		bag:        match[1],
		allowed:    make([]string, 0),
		quantities: make([]int, 0),
	}
	innerText := match[2]

	if innerText == "no other bags" {
		return rule, true
	}

	parts := strings.Split(innerText, ", ")
	for _, p := range parts {
		match = InnerRulePattern.FindStringSubmatch(p)
		if match == nil {
			return Rule{}, false
		}
		rule.allowed = append(rule.allowed, match[2])

		quantity, err := strconv.Atoi(match[1])
		if err != nil {
			return Rule{}, false
		}
		rule.quantities = append(rule.quantities, quantity)
	}

	return rule, true
}
