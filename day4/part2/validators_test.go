package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type validationTestCase struct {
	line     string
	expected bool
	fun      Validator
}

func TestValidationFunctions(t *testing.T) {
	cases := []validationTestCase{
		{"byr:2002", true, ValidateBirthYear},
		{"byr:2003", false, ValidateBirthYear},
		{"iyr:2010", true, ValidateIssueYear},
		{"iyr:2021", false, ValidateIssueYear},
		{"eyr:2020", true, ValidateExpirationYear},
		{"eyr:2031", false, ValidateExpirationYear},
		{"hgt:60in", true, ValidateHeight},
		{"hgt:190cm", true, ValidateHeight},
		{"hgt:190in", false, ValidateHeight},
		{"hgt:190", false, ValidateHeight},
		{"hcl:#123abc", true, ValidateHairColor},
		{"hcl:#123abz", false, ValidateHairColor},
		{"hcl:123abc", false, ValidateHairColor},
		{"ecl:brn", true, ValidateEyeColor},
		{"ecl:wat", false, ValidateEyeColor},
		{"pid:000000001", true, ValidatePassportID},
		{"pid:0123456789", false, ValidatePassportID},
	}

	for i, c := range cases {
		validation := c.fun(c.line)
		msg := fmt.Sprintf("case-%d: %s: expected %t (was %t)\n", i, c.line, c.expected, validation)
		assert.Equal(t, c.expected, validation, msg)
	}
}
