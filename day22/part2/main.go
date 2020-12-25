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

	_, score := PlayGame(cards[0], cards[1])
	fmt.Printf("solution code: %d\n", score)
}

func PlayGame(player1 Deck, player2 Deck) (int, int) {

	// sets to check for previous deck configuration
	player1Set := DeckSet{}
	player2Set := DeckSet{}

	// play rounds until cards empty
	for player1.NotEmpty() && player2.NotEmpty() {

		// if configuration was previously seen, player 1 wins instantly
		if player1Set.Contains(player1) || player2Set.Contains(player2) {
			return 1, player1.Score()
		}

		// save current deck configuration for later checks
		player1Set.Add(player1)
		player2Set.Add(player2)

		// draw cards
		player1Card := player1.Draw()
		player2Card := player2.Draw()

		var winner int

		// check if recursive game is necessary
		if player1Card <= player1.Size() && player2Card <= player2.Size() {
			player1Cutted := player1.CutCards(player1Card)
			player2Cutted := player2.CutCards(player2Card)
			winner, _ = PlayGame(player1Cutted, player2Cutted)
		} else if player1Card > player2Card {
			winner = 1
		} else {
			winner = 2
		}

		switch winner {
		case 1:
			player1.InsertAtBottom(player1Card, player2Card)
		case 2:
			player2.InsertAtBottom(player2Card, player1Card)
		}
	}

	if player1.NotEmpty() {
		return 1, player1.Score()
	} else {
		return 2, player2.Score()
	}
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
