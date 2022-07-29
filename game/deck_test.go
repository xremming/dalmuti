package game

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckHas80Cards(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, DeckSize, len(deck.cards))
}

func TestDeckContainsAtLeastOneOf(t *testing.T) {
	deck := NewDeck()

	cards_in_deck := make(map[Card]bool)

	for _, card := range deck.cards {
		cards_in_deck[card] = true
	}

	for _, card := range AllCards() {
		t.Run(card.String(), func(t *testing.T) {
			assert.True(t, cards_in_deck[card], "Expected deck to contain at least one %s", card)
		})
	}
}

func TestDeckHasCorrectAmountOfCards(t *testing.T) {
	deck := NewDeck()
	counts := make(map[Card]int)

	for _, card := range deck.cards {
		counts[card]++
	}

	for card, count := range counts {
		t.Run(fmt.Sprintf("%dx%s", card.CardsInDeck(), card), func(t *testing.T) {
			assert.Equal(t, card.CardsInDeck(), count, "Expected %d cards of type %s, got %d", card.CardsInDeck(), card, count)
		})
	}
}

func TestDeckCanDealCardsTo(t *testing.T) {
	for i := MinPlayers; i <= MaxPlayers; i++ {
		t.Run(fmt.Sprintf("%d players", i), func(t *testing.T) {
			deck := NewDeck()
			_, err := deck.Deal(i)
			if err != nil {
				t.Errorf("Error dealing %d players: %s", i, err)
			}
		})
	}
}

func TestDeckCannotDealCardsTo(t *testing.T) {
	for _, i := range []int{math.MinInt, -1, 0, 1, 11, 12, 13, math.MaxInt} {
		t.Run(fmt.Sprintf("%d players", i), func(t *testing.T) {
			deck := NewDeck()
			_, err := deck.Deal(i)
			if err == nil {
				t.Errorf("Expected error dealing %d players, got nil", i)
			}
		})
	}
}

func TestDeckDealsAllOfItsCardsWith(t *testing.T) {
	deck := NewDeck()
	for players := MinPlayers; players <= MaxPlayers; players++ {
		t.Run(fmt.Sprintf("%d players", players), func(t *testing.T) {
			hands, err := deck.Deal(players)
			if err != nil {
				t.Errorf("Error dealing %d players: %s", players, err)
			}

			if len(hands) != players {
				t.Errorf("Expected %d hands, got %d", players, len(hands))
			}

			total_cards_dealt := 0

			for _, hand := range hands {
				total_cards_dealt += len(hand.cards)
			}

			if total_cards_dealt != len(deck.cards) {
				t.Errorf("Expected %d cards dealt, got %d", len(deck.cards), total_cards_dealt)
			}
		})
	}
}

func TestDeckDealsExtraCardsToFirstPlayersWith(t *testing.T) {
	deck := NewDeck()
	for players := MinPlayers; players <= MaxPlayers; players++ {
		cards_per_player := DeckSize / players
		players_with_an_extra_card := DeckSize % players

		t.Run(fmt.Sprintf("%d players", players), func(t *testing.T) {
			hands, err := deck.Deal(players)
			if err != nil {
				t.Errorf("Error dealing %d players: %s", players, err)
			}

			if len(hands) != players {
				t.Errorf("Expected %d hands, got %d", players, len(hands))
			}

			for i, hand := range hands {
				if i < players_with_an_extra_card {
					// the first players_with_an_extra_card players should have one extra card
					if len(hand.cards) != cards_per_player+1 {
						t.Errorf("Expected an additional card to be dealt to player %d, got %d", i, len(hand.cards))
					}
				} else {
					// all other players should have cards_per_player cards
					if len(hand.cards) != cards_per_player {
						t.Errorf("Expected %d cards dealt to player %d, got %d", cards_per_player, i, len(hand.cards))
					}
				}
			}
		})
	}
}
