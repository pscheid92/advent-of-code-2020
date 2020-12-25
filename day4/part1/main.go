package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day4/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	batch := helpers.StackMultilineSeparatedByEmptyOne(lines)
	acceptor := NewAcceptor("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid") // without "cid"

	counter := 0
	for _, b := range batch {
		if acceptor.PresentSentence(b) {
			counter++
		}
	}

	fmt.Printf("%d of %d cards in batch are valid\n", counter, len(batch))
	fmt.Printf("solution code: %d\n", counter)
}

type Acceptor []*AcceptorData

type AcceptorData struct {
	index      int
	word       []rune
	len        int
	terminated bool
}

func NewAcceptor(words ...string) Acceptor {
	acceptor := make(Acceptor, len(words))

	for i, w := range words {
		runes := []rune(w)
		acceptor[i] = &AcceptorData{
			index:      0,
			word:       runes,
			len:        len(runes),
			terminated: false,
		}
	}

	return acceptor
}

func (a Acceptor) Present(r rune) {
	for _, d := range a {
		if d.terminated {
			continue
		}

		if d.word[d.index] != r {
			d.index = 0
			continue
		}

		d.index++
		if d.index >= d.len {
			d.terminated = true
		}
	}
}

func (a Acceptor) AllTerminated() bool {
	counter := 0
	for _, d := range a {
		if d.terminated {
			counter++
		}
	}
	return counter == len(a)
}

func (a Acceptor) Reset() {
	for _, d := range a {
		d.index = 0
		d.terminated = false
	}
}

func (a Acceptor) PresentSentence(sentence string) bool {
	a.Reset()
	for _, c := range sentence {
		a.Present(c)
	}
	return a.AllTerminated()
}
