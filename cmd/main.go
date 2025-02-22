package main

import (
	"fmt"
	"math"

	"github.com/levniko/okey-game/internal/game"
)

func main() {
	deck := game.NewDeck()

	fmt.Printf("Okey: %d\n", deck.GetOkey())
	fmt.Printf("Indicator: %d\n", deck.GetIndicator())

	players := deck.GetPlayers()

	for i, player := range players {
		fmt.Printf("%d. player:\n%v\n", i+1, player.GetHand())
	}

	scores := make([]int, 4)
	for i, player := range players {
		scores[i] = int(math.Max(
			float64(game.CalculateHandScore(player.GetHand())),
			float64(game.CalculatePairScore(player.GetHand())),
		))
	}

	fmt.Println("Scores:")
	winner := -1
	maxScore := -1
	for i, score := range scores {
		if score > maxScore {
			maxScore = score
			winner = i
		}
		fmt.Printf("%d. player: %d\n", i+1, score)
	}
	fmt.Printf("Best hand: Player %d \n", winner+1)
}
