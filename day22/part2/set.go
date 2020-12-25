package main

import (
	"strconv"
	"strings"
)

var empty = struct{}{}

type DeckSet map[string]struct{}

func (s DeckSet) Add(deck Deck) {
	key := convertDeckToKey(deck)
	s[key] = empty
}

func (s DeckSet) Contains(deck Deck) bool {
	key := convertDeckToKey(deck)
	_, found := s[key]
	return found
}

func convertDeckToKey(deck Deck) string {
	var builder strings.Builder
	for _, card := range deck {
		builder.WriteString(strconv.Itoa(card))
		builder.WriteRune(';')
	}
	return builder.String()
}
