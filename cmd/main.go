package main

import (
	"fmt"
	"strings"

	"github.com/levniko/okey-game/internal/game"
)

func main() {
	deck := game.NewDeck()
	players := deck.DistributeTiles()
	okeyTile := deck.GetOkeyTile()

	fmt.Println("\n=== OKEY DISTRIBUTION ===")
	fmt.Printf("Indicator: %v\n", deck.GetIndicator())
	fmt.Printf("Okey: %v\n\n", okeyTile)

	var bestPlayer *game.Player
	maxScore := -1

	for i := range players {
		hand := game.NewHand(players[i].Tiles, okeyTile)
		players[i].Score = hand.MaxScore
		players[i].Combinations = hand.Combinations

		var groups, series, pairs []game.Combination
		for _, comb := range hand.Combinations {
			switch comb.Type {
			case game.Group:
				groups = append(groups, comb)
			case game.Series:
				series = append(series, comb)
			case game.Pair:
				pairs = append(pairs, comb)
			}
		}

		fmt.Printf("Player %d Hand Analysis:\n", i+1)
		fmt.Println("▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀")
		printTiles(players[i].Tiles)

		fmt.Println("\nUsed Combinations:")
		printUsedCombinations(hand.MaxScore, groups, series, pairs)
		fmt.Printf("Total Score: %d\n", players[i].Score)
		fmt.Println(strings.Repeat("-", 80))

		if players[i].Score > maxScore {
			maxScore = players[i].Score
			bestPlayer = &players[i]
		}
	}

	fmt.Printf("\n🏆 WINNER: Player %d (%d points)\n",
		bestPlayer.ID+1,
		bestPlayer.Score)
}

func printUsedCombinations(score int, groups, series, pairs []game.Combination) {
	pairsScore := len(pairs) * 2
	if pairsScore > 0 && score == pairsScore {
		for _, p := range pairs {
			fmt.Printf("  ➜ Pair: %v\n", tilesToString(p.Tiles))
		}
		return
	}

	groupScore := 0
	for _, g := range groups {
		if len(g.Tiles) >= 3 {
			groupScore += len(g.Tiles)
		}
	}
	if groupScore > 0 && score == groupScore {
		for _, g := range groups {
			if len(g.Tiles) >= 3 {
				fmt.Printf("  ➜ Group (%d tiles): %v\n", len(g.Tiles), tilesToString(g.Tiles))
			}
		}
		return
	}

	seriesScore := 0
	for _, s := range series {
		if len(s.Tiles) >= 3 {
			seriesScore += len(s.Tiles)
		}
	}
	if seriesScore > 0 && score == seriesScore {
		for _, s := range series {
			if len(s.Tiles) >= 3 {
				fmt.Printf("  ➜ Series (%d tiles): %v\n", len(s.Tiles), tilesToString(s.Tiles))
			}
		}
		return
	}

	usedGroups := getUsedGroups(groups, score)
	usedSeries := getUsedSeries(series, score)
	usedPairs := getUsedPairs(pairs, score)

	if len(usedGroups) > 0 || len(usedSeries) > 0 || len(usedPairs) > 0 {
		for _, g := range usedGroups {
			fmt.Printf("  ➜ Group (%d tiles): %v\n", len(g.Tiles), tilesToString(g.Tiles))
		}
		for _, s := range usedSeries {
			fmt.Printf("  ➜ Series (%d tiles): %v\n", len(s.Tiles), tilesToString(s.Tiles))
		}
		for _, p := range usedPairs {
			fmt.Printf("  ➜ Pair: %v\n", tilesToString(p.Tiles))
		}
	} else {
		fmt.Println("  No usable combinations found")
	}
}

func getUsedPairs(pairs []game.Combination, totalScore int) []game.Combination {
	var used []game.Combination
	score := len(pairs) * 2
	if score > 0 && (totalScore == score || totalScore%score == 0) {
		used = pairs
	}
	return used
}

func getUsedGroups(groups []game.Combination, totalScore int) []game.Combination {
	var used []game.Combination
	for _, g := range groups {
		if len(g.Tiles) >= 3 {
			used = append(used, g)
		}
	}
	return used
}

func getUsedSeries(series []game.Combination, totalScore int) []game.Combination {
	var used []game.Combination
	for _, s := range series {
		if len(s.Tiles) >= 3 {
			used = append(used, s)
		}
	}
	return used
}

func printTiles(tiles []game.Tile) {
	colorMap := make(map[game.Color][]string)
	for _, t := range tiles {
		colorMap[t.Color] = append(colorMap[t.Color], t.String())
	}

	colors := []game.Color{game.Yellow, game.Blue, game.Black, game.Red, game.Fake}
	for _, color := range colors {
		if tiles, ok := colorMap[color]; ok {
			fmt.Printf("  %s: %s\n", color.String(), strings.Join(tiles, ", "))
		}
	}
}

func tilesToString(tiles []game.Tile) string {
	var s []string
	for _, t := range tiles {
		s = append(s, t.String())
	}
	return strings.Join(s, " + ")
}
