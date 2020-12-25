package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	// read and format input
	lines, err := helpers.ReadLineByLineFromFile("day4/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	batch := helpers.StackMultilineSeparatedByEmptyOne(lines)

	// prepare validators
	validators := []Validator{
		ValidateBirthYear,
		ValidateIssueYear,
		ValidateExpirationYear,
		ValidateHeight,
		ValidateHairColor,
		ValidateEyeColor,
		ValidatePassportID,
	}

	// check cards and count valid ones
	counter := 0
	for _, card := range batch {
		if ValidateCard(card, validators) {
			counter++
		}
	}

	// output
	fmt.Printf("%d of %d cards in batch are valid\n", counter, len(batch))
	fmt.Printf("solution code: %d\n", counter)
}

type Validator func(string) bool

func ValidateCard(card string, validators []Validator) bool {
	for _, v := range validators {
		if !v(card) {
			return false
		}
	}
	return true
}
