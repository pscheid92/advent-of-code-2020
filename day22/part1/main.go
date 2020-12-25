package main

import (
	"fmt"
	"github.com/pscheid92/advent-of-code-2020/helpers"
	"log"
)

func main() {
	lines, err := helpers.ReadLineByLineFromFile("day22/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	cards, err := ExtractCardsFromInput(lines)
	if err != nil {
		log.Fatalln(err)
	}

	winner, _ := PlayGame(cards[0], cards[1])
	fmt.Printf("solution code: %d\n", winner.Score())
}

func PlayGame(player1 Deck, player2 Deck) (Deck, Deck) {
	for player1.NotEmpty() && player2.NotEmpty() {
		player1Card := player1.Draw()
		player2Card := player2.Draw()

		if player1Card > player2Card {
			player1.InsertAtBottom(player1Card, player2Card)
			continue
		}

		if player2Card > player1Card {
			player2.InsertAtBottom(player2Card, player1Card)
			continue
		}
	}

	winner := player1
	other := player2

	if player2.NotEmpty() {
		winner = player2
		other = player1
	}
	return winner, other
}

func ExtractCardsFromInput(lines []string) ([2]Deck, error) {
	groups := helpers.GroupMultilineSeparatedByEmptyOne(lines)

	cardsPlayerOne, err := helpers.ConvertLinesToNumbers(groups[0][1:])
	if err != nil {
		return [2]Deck{}, fmt.Errorf("error parsing cards for player 1: %w", err)
	}

	cardsPlayerTwo, err := helpers.ConvertLinesToNumbers(groups[1][1:])
	if err != nil {
		return [2]Deck{}, fmt.Errorf("error parsing cards for player 2: %w", err)
	}

	cardsPlayerOne = reverseIntSlice(cardsPlayerOne)
	cardsPlayerTwo = reverseIntSlice(cardsPlayerTwo)
	return [2]Deck{cardsPlayerOne, cardsPlayerTwo}, nil
}

func reverseIntSlice(xs []int) []int {
	n := len(xs)
	result := make([]int, n)
	for i, x := range xs {
		result[n-i-1] = x
	}
	return result
}
