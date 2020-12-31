package main

import (
	"regexp"
	"strconv"
)

var (
	heightRegex = regexp.MustCompile(`hgt:([0-9]{2,3})(cm|in)`)
	hairColorRegex = regexp.MustCompile(`hcl:#([0-9a-f]{6})`)
	eyeColorRegex = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
	passportIDRegex = regexp.MustCompile(`pid:[0-9]{9}(\W|$)`)
)

func createYearValidator(field string, min int, max int) Validator {
	r := regexp.MustCompile(field + ":([0-9]{4})")

	return func(line string) bool {
		match := r.FindStringSubmatch(line)
		if match == nil {
			return false
		}

		year, err := strconv.Atoi(match[1])
		if err != nil {
			return false
		}

		return min <= year && year <= max
	}
}

var (
	ValidateBirthYear      = createYearValidator("byr", 1920, 2002)
	ValidateIssueYear      = createYearValidator("iyr", 2010, 2020)
	ValidateExpirationYear = createYearValidator("eyr", 2020, 2030)
)

func ValidateHeight(line string) bool {
	match := heightRegex.FindStringSubmatch(line)
	if len(match) <= 0 {
		return false
	}

	height, err := strconv.Atoi(match[1])
	if err != nil {
		return false
	}

	switch match[2] {
	case "cm":
		return 150 <= height && height <= 193
	case "in":
		return 59 <= height && height <= 76
	default:
		return false
	}
}

func ValidateHairColor(line string) bool {
	return hairColorRegex.MatchString(line)
}

func ValidateEyeColor(line string) bool {
	return eyeColorRegex.MatchString(line)
}

func ValidatePassportID(line string) bool {
	return passportIDRegex.MatchString(line)
}
