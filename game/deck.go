package game

import (
	"errors"
	"math/rand"
)

const (
	DeckSize   = 80
	MinPlayers = 2
	MaxPlayers = 10
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	cards := make([]Card, 0, 80)
	for i := 1; i <= 13; i++ {
		card := Card(i)
		count := card.CardsInDeck()
		for j := 0; j < count; j++ {
			cards = append(cards, card)
		}
	}
	return &Deck{cards}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Deal(players int) ([]PlayerHand, error) {
	if players < MinPlayers {
		return nil, errors.New("cannot deal to less than 2 players")
	}
	if players > MaxPlayers {
		return nil, errors.New("cannot deal to more than 10 players")
	}

	hands := make([][]Card, players)
	for i, card := range d.cards {
		hands[i%players] = append(hands[i%players], card)
	}

	playerHands := make([]PlayerHand, players)
	for i, hand := range hands {
		playerHands[i] = NewPlayerHand(hand)
	}

	return playerHands, nil
}
