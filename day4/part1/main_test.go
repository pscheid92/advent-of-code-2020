package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkThroughForest(t *testing.T) {
	var lines = []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
	expectedResults := []bool{true, false, true, false}

	// without "cid"
	acceptor := NewAcceptor("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
	batch := helpers.StackMultilineSeparatedByEmptyOne(lines)

	for i, card := range batch {
		accepted := acceptor.PresentSentence(card)
		assert.Equal(t, expectedResults[i], accepted, "case #%d: expected %t, but was %t", i, expectedResults[i], accepted)
	}

	assert.Len(t, batch, 4)
}
