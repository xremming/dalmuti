package game

// PlayerHand represents the cards that a player has in their hand.
type PlayerHand struct {
	cards []Card
}

func NewPlayerHand(cards []Card) PlayerHand {
	return PlayerHand{cards}
}

func (h PlayerHand) AllSets() []Set {
	cardsAndCounts := make(map[Card]int)
	for _, card := range h.cards {
		cardsAndCounts[card]++
	}

	jesters := cardsAndCounts[Jester]

	sets := make([]Set, 0)
	for card, count := range cardsAndCounts {
		for i := 1; i <= count; i++ {
			if card == Jester {
				sets = append(sets, Set{card, i, 0})
			} else {
				for j := 0; j <= jesters; j++ {
					sets = append(sets, Set{card, i + j, j})
				}
			}
		}
	}

	return sets
}

func (h PlayerHand) Size() int {
	return len(h.cards)
}

func (h *PlayerHand) IsEmpty() bool {
	return len(h.cards) == 0
}

func (h *PlayerHand) removeCards(cardToRemove Card, count int) bool {
	cardsRemoved := 0

	newCards := make([]Card, 0, len(h.cards))

	for _, card := range h.cards {
		if card == cardToRemove {
			cardsRemoved++
		} else {
			newCards = append(newCards, card)
		}
	}

	h.cards = newCards
	return cardsRemoved == count
}

func (h *PlayerHand) Play(s Set) bool {
	jesters_removed := h.removeCards(Jester, s.jesters)
	cards_removed := h.removeCards(s.card, s.count)

	return jesters_removed && cards_removed
}
