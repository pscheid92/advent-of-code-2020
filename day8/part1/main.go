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

	machine := NewMachine()
	machine.LoadProgram(program)

	// we await an already seen error
	// therefore we handle everything else as an fatal error
	err = machine.RunProgram()
	if !errors.Is(err, AlreadySeen) {
		log.Fatalln(err)
	}

	// solution code is current acc
	fmt.Printf("solution code: %d\n", machine.Accumulator)
}
