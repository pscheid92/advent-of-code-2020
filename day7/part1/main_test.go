package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractRule(t *testing.T) {
	cases := []struct {
		line    string
		ok      bool
		bag     string
		allowed []string
	}{
		{"something that is not a rule", false, "", nil},
		{"dotted black bags contain no other bags.", true, "dotted black", []string{}},
		{"bright white bags contain 1 shiny gold bag.", true, "bright white", []string{"shiny gold"}},
		{"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.", true, "shiny gold", []string{"dark olive", "vibrant plum"}},
		{"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.", true, "vibrant plum", []string{"faded blue", "dotted black"}},
	}

	for _, c := range cases {
		bag, allowed, ok := ExtractRule(c.line)
		assert.Equal(t, c.ok, ok)
		assert.Equal(t, c.bag, bag)
		assert.ElementsMatch(t, c.allowed, allowed)
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

	expected := map[string][]string{
		"light red":    {"bright white", "muted yellow"},
		"dark orange":  {"bright white", "muted yellow"},
		"bright white": {"shiny gold"},
		"muted yellow": {"shiny gold", "faded blue"},
		"shiny gold":   {"dark olive", "vibrant plum"},
		"dark olive":   {"faded blue", "dotted black"},
		"vibrant plum": {"faded blue", "dotted black"},
		"faded blue":   {},
		"dotted black": {},
	}

	table, ok := BuildRuleTable(lines)
	assert.True(t, ok)
	assert.Equal(t, expected, table)
}

func TestDeterminePossibleOuterBags(t *testing.T) {
	rules := map[string][]string{
		"light red":    {"bright white", "muted yellow"},
		"dark orange":  {"bright white", "muted yellow"},
		"bright white": {"shiny gold"},
		"muted yellow": {"shiny gold", "faded blue"},
		"shiny gold":   {"dark olive", "vibrant plum"},
		"dark olive":   {"faded blue", "dotted black"},
		"vibrant plum": {"faded blue", "dotted black"},
		"faded blue":   {},
		"dotted black": {},
	}
	expected := []string{"bright white", "muted yellow", "dark orange", "light red"}

	outerBags := DeterminePossibleOuterBags("shiny gold", rules)
	assert.Len(t, outerBags, 4)
	assert.ElementsMatch(t, expected, outerBags)
}
