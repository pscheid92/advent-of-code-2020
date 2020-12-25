package main

import (
	"errors"
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	// read code from file
	code, err := helpers.ReadLineByLineFromFile("day8/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	program, err := ParseCodeToProgram(code)
	if err != nil {
		log.Fatalln(err)
	}

	acc, err := SearchTerminatingProgram(program)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("solution code: %d\n", acc)
}

func SearchTerminatingProgram(program []Instruction) (int, error) {
	machine := NewMachine()

	for i, instruction := range program {
		switch instruction.command {
		case Jump:
			program[i].command = NoOp
		case NoOp:
			program[i].command = Jump
		default:
			// skip mutation if no jmp or nop
			continue
		}

		machine.LoadProgram(program)
		err := machine.RunProgram()

		// pass unexpected errors to caller
		if err != nil && !errors.Is(err, AlreadySeen) {
			return 0, err
		}

		// mutation led to endless loop, try next mutation
		// restore original instruction (in-place)
		if errors.Is(err, AlreadySeen) {
			program[i].command = instruction.command
			continue
		}

		// found terminating program, return accumulator
		return machine.Accumulator, nil
	}

	return 0, errors.New("found no terminating mutation of program")
}
