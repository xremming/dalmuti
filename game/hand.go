package game

import (
	"log"
	"math/rand"
)

type Hand struct {
	players        []PlayerHand
	sets           []Set
	skippedPlayers map[int]bool
}

func NewHand(players []PlayerHand) Hand {
	return Hand{players, nil, make(map[int]bool)}
}

func (h Hand) CanPlay(s Set) bool {
	if len(h.sets) == 0 {
		return true
	}
	return h.sets[len(h.sets)-1].CanBeFollowedBy(s)
}

func (h *Hand) Skip(playerIndex int) {
	h.skippedPlayers[playerIndex] = true
}

func (h *Hand) PlaySet(s Set, playerHand *PlayerHand) bool {
	if !h.CanPlay(s) {
		return false
	}

	h.sets = append(h.sets, s)
	playerHand.Play(s)

	return true
}

func (h Hand) PlayersWithCards() int {
	count := 0
	for _, playersHand := range h.players {
		if !playersHand.IsEmpty() {
			count++
		}
	}
	return count
}

func (h Hand) CardsLeft() int {
	count := 0
	for _, playersHand := range h.players {
		count += playersHand.Size()
	}
	return count
}

func (h Hand) Play(startingPlayer int) (Hand, int, []int) {
	if startingPlayer < 0 || startingPlayer >= len(h.players) {
		log.Panicf("starting player index %d is out of range", startingPlayer)
	}

	currentWinner := -1
	playersOut := make([]int, 0)

	playerIndex := startingPlayer

	nextPlayer := func() int {
		playerIndex = (playerIndex + 1) % len(h.players)
		return playerIndex
	}

	playersSkipping := func() int {
		count := 0
		for _, skipped := range h.skippedPlayers {
			if skipped {
				count++
			}
		}
		return count
	}

	for currentPlayersHand := h.players[playerIndex]; playersSkipping() != len(h.players); currentPlayersHand = h.players[nextPlayer()] {
		// players with no cards in hand must always skip
		if currentPlayersHand.IsEmpty() {
			h.skippedPlayers[playerIndex] = true
			continue
		}

		// find all sets the current player can play
		allPlayerSets := currentPlayersHand.AllSets()
		playableSets := make([]Set, 0, len(allPlayerSets))
		for _, set := range allPlayerSets {
			if h.CanPlay(set) {
				playableSets = append(playableSets, set)
			}
		}

		// if the player cannot play a set, they must skip
		if len(playableSets) == 0 {
			h.skippedPlayers[playerIndex] = true
			continue
		}

		// select a random set to play
		selectedSet := playableSets[rand.Intn(len(playableSets))]
		h.PlaySet(selectedSet, &h.players[playerIndex])

		// player played a set, so they are now winning
		currentWinner = playerIndex

		// player played a set, so they are no longer skipped
		h.skippedPlayers[startingPlayer] = false

		// if the player no longer has cards in hand, they are out
		if h.players[playerIndex].IsEmpty() {
			playersOut = append(playersOut, playerIndex)
		}
	}

	if currentWinner == -1 {
		log.Panicf("no winner found")
	}

	return Hand{h.players, nil, make(map[int]bool)}, currentWinner, playersOut
}
