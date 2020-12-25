package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ConvertLinesToNumbers(lines []string) ([]int, error) {
	var numbers []int

	for _, l := range lines {
		x, err := strconv.Atoi(l)
		if err != nil {
			return nil, fmt.Errorf("error converting '%s' to number", l)
		}
		numbers = append(numbers, x)
	}

	return numbers, nil
}

func ReadLineByLine(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}
	return lines
}

func ReadLineByLineFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", filename, err)
	}
	defer file.Close()

	lines := ReadLineByLine(file)
	return lines, nil
}

func StackMultilineSeparatedByEmptyOne(lines []string) []string {
	batches := make([]string, 0)

	var currentLine string
	for _, l := range lines {
		// use empty line as marker
		if l == "" {
			batches = append(batches, currentLine)
			currentLine = ""
			continue
		}

		currentLine += " "
		currentLine += l
	}

	// handle last line (missing empty line at the end of input)
	if currentLine != "" {
		batches = append(batches, currentLine)
	}

	return batches
}

func GroupMultilineSeparatedByEmptyOne(lines []string) [][]string {
	groups := make([][]string, 0)

	current := make([]string, 0)
	for _, l := range lines {
		// use empty line as marker
		if l == "" {
			groups = append(groups, current)
			current = make([]string, 0)
			continue
		}

		current = append(current, l)
	}

	// handle last line (missing empty line at the end of input)
	if len(current) != 0 {
		groups = append(groups, current)
	}

	return groups
}
