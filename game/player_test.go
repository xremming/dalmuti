package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandAllSetsSingleCard(t *testing.T) {
	for _, card := range AllCards() {
		t.Run(card.String(), func(t *testing.T) {
			hand := NewPlayerHand([]Card{card})
			sets := hand.AllSets()

			if len(sets) != 1 {
				t.Errorf("Expected 1 set, got %d", len(sets))
			}

			assert.Equal(t, Set{card, 1, 0}, sets[0])
		})
	}
}

func TestHandAllSetsSingleCardAndJester(t *testing.T) {
	for _, card := range AllCards() {
		t.Run(card.String(), func(t *testing.T) {
			if card == Jester {
				t.SkipNow()
			}

			hand := NewPlayerHand([]Card{card, Jester})
			sets := hand.AllSets()

			if len(sets) != 3 {
				t.Errorf("Expected 3 sets, got %d", len(sets))
			}

			expected := []Set{
				{card, 1, 0},
				{card, 2, 1},
				{Jester, 1, 0},
			}

			assert.ElementsMatch(t, expected, sets)
		})
	}
}

func TestAllSetsTwoJesters(t *testing.T) {
	hand := NewPlayerHand([]Card{Jester, Jester})
	sets := hand.AllSets()
	expected := []Set{
		{Jester, 1, 0},
		{Jester, 2, 0},
	}
	assert.ElementsMatch(t, expected, sets)
}
