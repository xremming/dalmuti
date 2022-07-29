package game

import (
	"log"
)

func PlayRound(players []string) []int {
	winners := make([]int, 0, len(players))

	totalPlayers := len(players)

	deck := NewDeck()
	deck.Shuffle()

	hands, err := deck.Deal(totalPlayers)
	if err != nil {
		log.Panic(err)
	}

	for {
		playersOut := 0
		for _, hand := range hands {
			if hand.IsEmpty() {
				playersOut++
			}
		}

		cardsLeft := 0
		for _, hand := range hands {
			cardsLeft += hand.Size()
		}

		log.Printf("%d players out, %d cards left", playersOut, cardsLeft)

		if playersOut == len(hands)-1 {
			break
		}

		hand := NewHand(hands)
		log.Printf("new hand")
		_ = hand
	}

	for i, hand := range hands {
		if !hand.IsEmpty() {
			winners = append(winners, i)
		}
	}

	return winners
}
