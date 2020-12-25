package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	expressions, err := helpers.ReadLineByLineFromFile("day18/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sum := uint64(0)
	for _, expression := range expressions {
		result, err := Evaluate(expression)
		if err != nil {
			log.Fatalln(err)
		}
		sum += uint64(result)
	}

	fmt.Printf("solution code: %d\n", sum)
}

// Evaluate evaluates a given mathematical expression
func Evaluate(expression string) (int, error) {
	rpnTokens, err := preprocessExpression(expression)
	if err != nil {
		return 0, fmt.Errorf("error evaluating expression (phase 1): %w", err)
	}

	result, err := processRPN(rpnTokens)
	if err != nil {
		return 0, fmt.Errorf("error evaluating expression (phase 2): %w", err)
	}

	return result, nil
}

// preprocessExpression uses the shunting-yard algorithm to convert from infix to reverse polish notation
func preprocessExpression(expression string) ([]string, error) {
	output := StringStack{}
	operations := StringStack{}

	var s scanner.Scanner
	s.Init(strings.NewReader(expression))

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		word := s.TokenText()

		if _, err := strconv.Atoi(word); err == nil {
			output.Push(word)
			continue
		}

		// check if opening bracket
		if word == "(" {
			operations.Push(word)
			continue
		}

		// check if closing bracket
		if word == ")" {
			for {
				element, err := operations.Pop()
				if err != nil {
					return nil, err
				}
				if element == "(" {
					break
				}
				output.Push(element)
			}
			continue
		}

		// check for operators
		if word == "+" || word == "*" {
			// clear the operator stack
			for {
				// as long as operation stack is not empty
				element, err := operations.Peek()
				if err != nil {
					break
				}

				// as long as operation stack head is an operator
				if element != "+" && element != "*" {
					break
				}

				// stop if precedence of operator is bigger than operator stack head
				if word == "+" && element == "*" {
					break
				}

				// take operator and push it to output
				element, err = operations.Pop()
				if err != nil {
					return nil, err
				}
				output.Push(element)
			}

			// push new operation to operator stack
			operations.Push(word)
		}
	}

	for len(operations) > 0 {
		element, err := operations.Pop()
		if err != nil {
			continue
		}

		if element == "(" {
			return nil, fmt.Errorf("more opening than closing brackets in '%s'", expression)
		}

		output.Push(element)
	}

	return output, nil
}

// processRPN takes a tokenized expression in reverse polish notation and evaluates it
func processRPN(tokens []string) (int, error) {
	accumulator := IntStack{}
	for _, op := range tokens {
		// is a number?
		if number, err := strconv.Atoi(op); err == nil {
			accumulator.Push(number)
			continue
		}

		// try to get first operand
		x, err := accumulator.Pop()
		if err != nil {
			return 0, fmt.Errorf("cannot processRPN operation '%s': missing first operand: %w", op, err)
		}

		// try to get second operand
		y, err := accumulator.Pop()
		if err != nil {
			return 0, fmt.Errorf("cannot processRPN operation '%s': missing second operand: %w", op, err)
		}

		// processRPN expression
		switch op {
		case "+":
			accumulator.Push(x + y)
		case "*":
			accumulator.Push(x * y)
		default:
			return 0, fmt.Errorf("unevaluable operand found '%s'", op)
		}
	}

	if len(accumulator) != 1 {
		return 0, fmt.Errorf("expected single remaining operand in accumulator, but found %d", len(accumulator))
	}
	return accumulator.Pop()
}
