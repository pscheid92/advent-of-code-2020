package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strings"
)

func main() {
	program, err := helpers.ReadLineByLineFromFile("day14/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	memory := NewMemory()
	err = RunProgram(memory, program)
	if err != nil {
		log.Fatalln(err)
	}

	sum := memory.SumData()
	fmt.Printf("solution code: %d\n", sum)
}

func RunProgram(memory Memory, program []string) error {
	// each line is an instruction
	for _, instruction := range program {

		// set new mask
		if strings.HasPrefix(instruction, "mask") {
			var maskSpecification string
			n, err := fmt.Sscanf(instruction, "mask = %s", &maskSpecification)
			if err != nil {
				return err
			}
			if n != 1 {
				return fmt.Errorf("tried to parse mask instruction, but found %d values instead of one", n)
			}
			memory.Mask = NewMask(maskSpecification)
			continue
		}

		// set memory
		if strings.HasPrefix(instruction, "mem[") {
			var address uint64
			var value uint64
			n, err := fmt.Sscanf(instruction, "mem[%d] = %d", &address, &value)
			if err != nil {
				return err
			}
			if n != 2 {
				return fmt.Errorf("tried to parse memory set instruction, but found %d values instead of two", n)
			}
			memory.Set(address, value)
			continue
		}
	}

	return nil
}

