package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"regexp"
	"strconv"
	"strings"
)

func ExtractInput(lines []string) ([]FieldSpecification, Ticket, []Ticket, error) {
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)

	ruleLines := groups[0]
	myTicketLine := groups[1][1]
	nearbyTicketLines := groups[2][1:]

	fieldSpecifications, err := ReadFieldSpecifications(ruleLines)
	if err != nil {
		return nil, nil, nil, err
	}

	myTicket, err := ReadMyTicket(myTicketLine)
	if err != nil {
		return nil, nil, nil, err
	}

	nearbyTickets, err := ReadNearbyTickets(nearbyTicketLines)
	if err != nil {
		return nil, nil, nil, err
	}

	return fieldSpecifications, myTicket, nearbyTickets, nil
}

func ReadMyTicket(line string) (Ticket, error) {
	fields := strings.Split(line, ",")
	numbers, err := helpers.ConvertLinesToNumbers(fields)
	return numbers, err
}

func ReadNearbyTickets(lines []string) ([]Ticket, error) {
	tickets := make([]Ticket, len(lines))

	for i, l := range lines {
		ticket, err := ReadMyTicket(l)
		if err != nil {
			return nil, err
		}
		tickets[i] = ticket
	}

	return tickets, nil
}

func ReadFieldSpecifications(lines []string) ([]FieldSpecification, error) {
	regex := regexp.MustCompile(`^(.*?): (\d+)-(\d+) or (\d+)-(\d+)$`)

	fieldSpecifications := make([]FieldSpecification, len(lines))
	for i, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("could not extract data in line '%s'", line)
		}

		field := matches[1]
		numbers := make([]int, 4)
		for j, m := range matches[2:6] {
			n, err := strconv.Atoi(m)
			if err != nil {
				return nil, err
			}
			numbers[j] = n
		}

		fieldSpecifications[i].Name = field
		fieldSpecifications[i].ranges[0][0] = numbers[0]
		fieldSpecifications[i].ranges[0][1] = numbers[1]
		fieldSpecifications[i].ranges[1][0] = numbers[2]
		fieldSpecifications[i].ranges[1][1] = numbers[3]
	}
	return fieldSpecifications, nil
}
