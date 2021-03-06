package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldSpecification_IsValid(t *testing.T) {
	rule := FieldSpecification{
		Name: "dummy",
		ranges: [2][2]int{
			{-3, 5},
			{10, 50},
		},
	}

	cases := []struct {
		input    int
		expected bool
	}{
		{-50, false},
		{0, true},
		{7, false},
		{20, true},
		{100, false},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, rule.IsValid(c.input))
	}
}

func TestCheckTicket(t *testing.T) {
	specifications := []FieldSpecification{
		{"class", [2][2]int{{1, 3}, {5, 7}}},
		{"row", [2][2]int{{6, 11}, {33, 44}}},
		{"seat", [2][2]int{{13, 40}, {45, 50}}},
	}

	cases := []struct {
		ticket        Ticket
		isValid       bool
		invalidColumn int
	}{
		{Ticket{7, 3, 47}, true, -1},
		{Ticket{40, 4, 50}, false, 4},
		{Ticket{55, 2, 20}, false, 55},
		{Ticket{38, 6, 12}, false, 12},
	}

	for _, c := range cases {
		isValid, invalidColumn := CheckTicket(c.ticket, specifications)

		// invalid test cases
		if !c.isValid {
			assert.False(t, isValid)
			assert.Equal(t, c.invalidColumn, invalidColumn)
			continue
		}

		// valid test cases
		assert.True(t, isValid)
	}
}
