package main

type Deck []int

func NewDeck(cards ...int) Deck {
	return cards
}

func (d *Deck) Draw() int {
	end := len(*d) - 1
	result := (*d)[end]
	*d = (*d)[:end]
	return result
}

func (d *Deck) InsertAtBottom(cards ...int) {
	cards = reverseIntSlice(cards)
	*d = append(cards, *d...)
}

func (d Deck) Empty() bool {
	return len(d) == 0
}

func (d Deck) NotEmpty() bool {
	return !d.Empty()
}

func (d Deck) Score() int {
	score := 0
	for i, c := range d {
		score += (i + 1) * c
	}
	return score
}
