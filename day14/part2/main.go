package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strings"
)

func main() {
	// load program
	program, err := helpers.ReadLineByLineFromFile("day14/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// create memory and run program on it
	memory := make(map[uint64]uint64)
	err = RunProgram(memory, program)
	if err != nil {
		log.Fatalln(err)
	}

	// summarize data and print solution code
	sum := SumValues(memory)
	fmt.Printf("solution code: %d\n", sum)
}

func RunProgram(memory map[uint64]uint64, program []string) error {
	var currentMask string

	// each line is an instruction
	for _, instruction := range program {

		// consider mask instruction
		if mask, found, err := tryScanMask(instruction); found {
			if err != nil {
				return err
			}
			currentMask = mask
		}

		// consider memory instruction
		if addr, value, found, err := tryScanMemory(instruction); found {
			if err != nil {
				return err
			}

			for _, a := range CalculateAddresses(currentMask, addr) {
				memory[a] = value
			}
		}
	}

	return nil
}

func tryScanMask(instruction string) (string, bool, error) {
	if !strings.HasPrefix(instruction, "mask") {
		return "", false, nil
	}

	var mask string
	n, err := fmt.Sscanf(instruction, "mask = %s", &mask)
	if err != nil {
		return "", true, err
	}
	if n != 1 {
		err = fmt.Errorf("tried to parse memory set instruction, but found %d values instead of two", n)
		return "", true, err
	}

	return mask, true, nil
}

func tryScanMemory(instruction string) (uint64, uint64, bool, error) {
	if !strings.HasPrefix(instruction, "mem[") {
		return 0, 0, false, nil
	}

	var address uint64
	var value uint64
	n, err := fmt.Sscanf(instruction, "mem[%d] = %d", &address, &value)
	if err != nil {
		return 0, 0, true, err
	}
	if n != 2 {
		err = fmt.Errorf("tried to parse memory set instruction, but found %d values instead of two", n)
		return 0, 0, true, err
	}

	return address, value, true, nil
}

func CalculateAddresses(maskSpecification string, address uint64) []uint64 {
	results := []uint64{address}
	bitMarker := uint64(1) << 35

	for _, m := range maskSpecification {
		// set to 1
		if m == '1' {
			for i := 0; i < len(results); i++ {
				results[i] |= bitMarker
			}
		}

		// duplicate and set one version to 0 and the other to 1
		if m == 'X' {
			n := len(results)
			for i := 0; i < n; i++ {
				bitSet := results[i] | bitMarker
				bitCleared := results[i] & ^bitMarker

				results[i] = bitCleared
				results = append(results, bitSet)
			}
		}

		// move marker to next position
		bitMarker >>= 1
	}

	return results
}

func SumValues(memory map[uint64]uint64) uint64 {
	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}
