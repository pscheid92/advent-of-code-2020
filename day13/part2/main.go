package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day13/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	busses, err := ParseInput(lines[1])
	if err != nil {
		log.Fatalln(err)
	}

	ns, bs := PrepareInputForCRT(busses)
	timestamp := ChineseReminderTheorem(ns, bs)
	fmt.Printf("solution code: %d\n", timestamp)
}

func ParseInput(line string) (map[int]int, error) {
	busses := make(map[int]int)
	for i, b := range strings.Split(line, ",") {
		// ignore x entries
		if b == "x" {
			continue
		}

		// parse the bus id
		id, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}

		// save bus (i is time delay)
		busses[i] = id
	}

	return busses, nil
}

// PrepareInputForCRT takes the busses map and derives input for CRT from it
func PrepareInputForCRT(busses map[int]int) ([]uint64, []uint64) {
	ns := make([]uint64, 0, len(busses))
	bs := make([]uint64, 0, len(busses))
	for delay, id := range busses {
		ns = append(ns, uint64(id))
		bs = append(bs, uint64(id-delay))
	}
	return ns, bs
}

// ChineseReminderTheorem is a naive implementation of the chinese remainder theorem
func ChineseReminderTheorem(ns []uint64, bs []uint64) uint64 {
	N := uint64(1)
	for _, n := range ns {
		N *= n
	}

	sum := uint64(0)
	for i := 0; i < len(ns); i++ {
		n := N / ns[i]
		x := CalculateModularInverse(n, ns[i])
		sum += bs[i] * n * x
	}

	return sum % N
}

// CalculateModularInverse searches for the modular inverse by trial and error.
func CalculateModularInverse(number uint64, divisor uint64) uint64 {
	var inverse uint64
	for inverse = uint64(0); inverse < number-1; inverse++ {
		if (number*inverse)%divisor == 1 {
			break
		}
	}
	return inverse
}
