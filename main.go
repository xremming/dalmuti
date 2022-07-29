package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/polarpayne/dalmuti/game"
	"github.com/schollz/progressbar/v3"
)

var (
	it      int
	players int
)

func init() {
	flag.IntVar(&it, "it", 10_000, "number of iterations")
	flag.IntVar(&players, "players", 4, "number of players")

	flag.Parse()
}

func main() {
	winners := make([]int, 0, it)
	losers := make([]int, 0, it)

	bar := progressbar.NewOptions(it,
		progressbar.OptionShowIts(),
		progressbar.OptionThrottle(1*time.Second),
	)
	defer bar.Close()

	for i := 0; i < it; i++ {
		deck := game.NewDeck()
		deck.Shuffle()

		playerHands, err := deck.Deal(players)
		if err != nil {
			log.Panicf("error dealing hands: %s", err)
		}

		hand := game.NewHand(playerHands)
		newOrder := make([]int, 0, players)

		hand, winner, playersOut := hand.Play(0)
		newOrder = append(newOrder, playersOut...)

		for {
			if hand.PlayersWithCards() == 0 {
				break
			}

			hand, winner, playersOut = hand.Play(winner)
			newOrder = append(newOrder, playersOut...)
		}

		winners = append(winners, newOrder[0])
		losers = append(losers, newOrder[len(newOrder)-1])

		bar.Add(1)
	}
	bar.Finish()

	winnerCounts := make(map[int]int)
	for _, winner := range winners {
		winnerCounts[winner]++
	}

	loserCounts := make(map[int]int)
	for _, loser := range losers {
		loserCounts[loser]++
	}

	fmt.Print("\nresults:\n")

	for i := 0; i < players; i++ {
		fmt.Printf("\tpos %d won  %5.2f%% of the time\n", i, 100.0*float64(winnerCounts[i])/float64(it))
		fmt.Printf("\tpos %d lost %5.2f%% of the time\n", i, 100.0*float64(loserCounts[i])/float64(it))
	}

	// spew.Dump(winnerCounts, loserCounts)
}
