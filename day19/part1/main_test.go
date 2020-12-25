package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRules(t *testing.T) {
	lines := []string{
		"0: 1 2",
		"1: \"a\"",
		"2: 1 3 | 3 1",
		"3: \"b\"",
	}

	expectedLookup := map[string][]string{
		"1,2": {"0"},
		"1,3": {"2"},
		"3,1": {"2"},
		"a":   {"1"},
		"b":   {"3"},
	}

	// should have discovered 5 rules
	lookup := ParseRules(lines)
	assert.Len(t, lookup, 5)

	// compare with matching entry in `expectedLookup`
	// remove entry after match
	for key, set := range lookup {
		elements := set.GetAll()
		assert.ElementsMatch(t, expectedLookup[key], elements)
		delete(expectedLookup, key)
	}

	// all entries should have been matched
	assert.Len(t, expectedLookup, 0)
}

func TestIsValidWord(t *testing.T) {
	// rule 0 is manually converted to Chomsky Normal Form (CNF) by
	// replacing `0: 4 1 5` by rules `0: 6 5` and `6: 4 1`
	lines := []string{
		"0: 6 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		"4: \"a\"",
		"5: \"b\"",
		"6: 4 1",
	}

	cases := []struct {
		word  string
		valid bool
	}{
		{"ababbb", true},
		{"bababa", false},
		{"abbbab", true},
		{"aaabbb", false},
		{"aaaabbb", false},
	}

	lookup := ParseRules(lines)
	for _, c := range cases {
		valid := IsValidWord(c.word, lookup, "0")
		assert.Equal(t, c.valid, valid)
	}
}
