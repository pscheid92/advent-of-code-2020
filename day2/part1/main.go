package main

import (
	"errors"
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	// read lines
	lines, err := helpers.ReadLineByLineFromFile("day2/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// convert lines to problems
	problems := make([]ProblemInstance, 0)
	for _, l := range lines {
		p, err := NewProblemInstanceFromLine(l)
		if err != nil {
			log.Fatalln(err)
		}

		problems = append(problems, p)
	}

	// count valid problems
	// overall complexity O(n * m) with n = len(problems) and m = length of longest password
	successCount := CountValidProblems(problems)

	// final output
	fmt.Printf("%d passwords are valid\n", successCount)
	fmt.Printf("answer code is %d\n", successCount)
}

func NewProblemInstanceFromLine(line string) (ProblemInstance, error) {
	var p ProblemInstance

	n, err := fmt.Sscanf(line, "%d-%d %c: %s\n", &p.min, &p.max, &p.char, &p.password)
	if err != nil {
		return ProblemInstance{}, err
	}
	if n < 4 {
		msg := fmt.Sprintf("expected 4 errors, got only %d from line '%s'\n", n, line)
		return ProblemInstance{}, errors.New(msg)
	}

	return p, nil
}

func CountValidProblems(problems []ProblemInstance) int {
	// complexity is O(n) where n = number of problem instances

	counter := 0
	for _, p := range problems {
		if p.IsValidPassword() {
			counter++
		}
	}
	return counter
}

type ProblemInstance struct {
	min      int
	max      int
	char     rune
	password string
}

func (p ProblemInstance) IsValidPassword() bool {
	// complexity O(n) with n = length of password

	counter := 0
	for _, c := range p.password {
		if p.char == c {
			counter++
		}
	}

	return p.min <= counter && counter <= p.max
}
