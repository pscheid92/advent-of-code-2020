package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractRule(t *testing.T) {
	cases := []struct {
		line string
		ok   bool
		rule Rule
	}{
		{"something that is not a rule", false, Rule{}},
		{"dotted black bags contain no other bags.", true, Rule{"dotted black", []string{}, []int{}}},
		{"bright white bags contain 1 shiny gold bag.", true, Rule{"bright white", []string{"shiny gold"}, []int{1}}},
		{"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.", true, Rule{"shiny gold", []string{"dark olive", "vibrant plum"}, []int{1, 2}}},
		{"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.", true, Rule{"vibrant plum", []string{"faded blue", "dotted black"}, []int{5, 6}}},
	}

	for _, c := range cases {
		rule, ok := ExtractRule(c.line)
		assert.Equal(t, c.ok, ok)
		assert.Equal(t, c.rule.bag, rule.bag)
		assert.ElementsMatch(t, c.rule.allowed, rule.allowed)
		assert.ElementsMatch(t, c.rule.quantities, rule.quantities)
	}
}

func TestBuildRuleTable(t *testing.T) {
	var lines = []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	expected := map[string]Rule{
		"light red":    {"light red", []string{"bright white", "muted yellow"}, []int{1, 2}},
		"dark orange":  {"dark orange", []string{"bright white", "muted yellow"}, []int{3, 4}},
		"bright white": {"bright white", []string{"shiny gold"}, []int{1}},
		"muted yellow": {"muted yellow", []string{"shiny gold", "faded blue"}, []int{2, 9}},
		"shiny gold":   {"shiny gold", []string{"dark olive", "vibrant plum"}, []int{1, 2}},
		"dark olive":   {"dark olive", []string{"faded blue", "dotted black"}, []int{3, 4}},
		"vibrant plum": {"vibrant plum", []string{"faded blue", "dotted black"}, []int{5, 6}},
		"faded blue":   {"faded blue", []string{}, []int{}},
		"dotted black": {"dotted black", []string{}, []int{}},
	}

	table, ok := BuildRuleTable(lines)
	assert.True(t, ok)
	assert.Equal(t, expected, table)
}

func TestCalculateContainedBags(t *testing.T) {
	rules := map[string]Rule{
		"light red":    {"light red", []string{"bright white", "muted yellow"}, []int{1, 2}},
		"dark orange":  {"dark orange", []string{"bright white", "muted yellow"}, []int{3, 4}},
		"bright white": {"bright white", []string{"shiny gold"}, []int{1}},
		"muted yellow": {"muted yellow", []string{"shiny gold", "faded blue"}, []int{2, 9}},
		"shiny gold":   {"shiny gold", []string{"dark olive", "vibrant plum"}, []int{1, 2}},
		"dark olive":   {"dark olive", []string{"faded blue", "dotted black"}, []int{3, 4}},
		"vibrant plum": {"vibrant plum", []string{"faded blue", "dotted black"}, []int{5, 6}},
		"faded blue":   {"faded blue", []string{}, []int{}},
		"dotted black": {"dotted black", []string{}, []int{}},
	}
	count := CalculateContainedBags("shiny gold", rules)
	assert.Equal(t, 32, count)

	rules = map[string]Rule{
		"shiny gold":  {bag: "shiny gold", allowed: []string{"dark red"}, quantities: []int{2}},
		"dark red":    {bag: "dark red", allowed: []string{"dark orange"}, quantities: []int{2}},
		"dark orange": {bag: "dark orange", allowed: []string{"dark yellow"}, quantities: []int{2}},
		"dark yellow": {bag: "dark yellow", allowed: []string{"dark green"}, quantities: []int{2}},
		"dark green":  {bag: "dark green", allowed: []string{"dark blue"}, quantities: []int{2}},
		"dark blue":   {bag: "dark blue", allowed: []string{"dark violet"}, quantities: []int{2}},
		"dark violet": {bag: "dark violet", allowed: []string{}, quantities: []int{}},
	}
	count = CalculateContainedBags("shiny gold", rules)
	assert.Equal(t, 126, count)
}
