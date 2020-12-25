package main

import (
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const input = `
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`

func TestExtractInput(t *testing.T) {
	expectedFieldSpecifications := []FieldSpecification{
		{"class", [2][2]int{{1, 3}, {5, 7}}},
		{"row", [2][2]int{{6, 11}, {33, 44}}},
		{"seat", [2][2]int{{13, 40}, {45, 50}}},
	}
	expectedMyTicket := Ticket{7, 1, 14}
	expectedNearbyTickets := []Ticket{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	reader := strings.NewReader(input[1:])
	lines := helpers.ReadLineByLine(reader)
	fieldSpecifications, myTicket, nearbyTickets, err := ExtractInput(lines)

	assert.NoError(t, err)
	assert.Equal(t, expectedMyTicket, myTicket)
	assert.ElementsMatch(t, expectedFieldSpecifications, fieldSpecifications)
	assert.ElementsMatch(t, expectedNearbyTickets, nearbyTickets)
}
