package game

import (
	"math"
	"sort"
)

func CalculateHandScore(hand []int) int {
	sortedHand := preprocessHand(hand)
	series := findAllSeries(sortedHand)
	groups := findAllGroups(sortedHand)
	pairs := findAllPairs(sortedHand) // Çiftleri ekle
	allCombos := mergeCombinations(series, groups)
	allCombos = mergeCombinations(allCombos, pairs)

	sort.Slice(allCombos, func(i, j int) bool {
		return len(allCombos[i]) > len(allCombos[j])
	})

	maxScore := backtrackCombinations(allCombos, 0, 0, make(map[int]bool))
	return int(math.Min(float64(maxScore), 14))
}

func mergeCombinations(a, b [][]int) [][]int {
	result := make([][]int, len(a))
	copy(result, a)
	return append(result, b...)
}

func backtrackCombinations(combos [][]int, index int, currentScore int, used map[int]bool) int {
	if index >= len(combos) {
		return currentScore
	}

	maxScore := backtrackCombinations(combos, index+1, currentScore, used)

	combo := combos[index]
	if canUseCombo(combo, used) {
		newUsed := make(map[int]bool)
		for k, v := range used {
			newUsed[k] = v
		}
		markUsed(combo, newUsed)

		scoreWithCombo := backtrackCombinations(combos, index+1, currentScore+len(combo), newUsed)
		if scoreWithCombo > maxScore {
			maxScore = scoreWithCombo
		}
	}

	return maxScore
}

func canUseCombo(combo []int, used map[int]bool) bool {
	for _, tile := range combo {
		if tile != OK && used[tile] {
			return false
		}
	}
	return true
}

func markUsed(combo []int, used map[int]bool) {
	for _, tile := range combo {
		if tile != OK {
			used[tile] = true
		}
	}
}

func preprocessHand(hand []int) []int {
	sorted := make([]int, 0)
	okeys := make([]int, 0)

	for _, tile := range hand {
		if tile == OK {
			okeys = append(okeys, tile)
		} else {
			sorted = append(sorted, tile)
		}
	}
	sort.Ints(sorted)
	return append(sorted, okeys...)
}

func findAllPairs(hand []int) [][]int {
	pairs := make([][]int, 0)
	counts := make(map[int]int)

	for _, tile := range hand {
		counts[tile]++
	}

	for tile, count := range counts {
		if count >= 2 {
			pairs = append(pairs, []int{tile, tile})
		}
	}

	okeyCount := counts[OK]
	if okeyCount >= 2 {
		pairs = append(pairs, []int{OK, OK})
	}

	return pairs
}

func CalculatePairScore(hand []int) int {
	pairs := 0
	for i := 1; i < len(hand); i++ {
		if hand[i] == hand[i-1] {
			pairs++
		}
	}
	if hand[len(hand)-1] == OK {
		pairs++
	}
	return pairs * 2
}

func findAllGroups(hand []int) [][]int {
	groups := make(map[int][]int)

	for _, tile := range hand {
		num := tile % 13
		groups[num] = append(groups[num], tile)
	}

	var validGroups [][]int
	for _, tiles := range groups {
		colorMap := make(map[int]bool)
		uniqueTiles := make([]int, 0)

		// Farklı renkleri seç
		for _, tile := range tiles {
			color := tile / 13
			if !colorMap[color] {
				colorMap[color] = true
				uniqueTiles = append(uniqueTiles, tile)
			}
		}

		if len(uniqueTiles) >= 3 {
			validGroups = append(validGroups, uniqueTiles[:3])
		}
	}

	return validGroups
}

func findAllSeries(hand []int) [][]int {
	var series [][]int
	for i := 0; i < len(hand); i++ {
		color := hand[i] / 13
		current := []int{hand[i]}

		for j := i + 1; j < len(hand); j++ {
			if hand[j]/13 != color {
				break
			}

			if hand[j]-1 == hand[j-1] {
				current = append(current, hand[j])
			} else if hand[j]-2 == hand[j-1] && hand[len(hand)-1] == OK {
				current = append(current, OK, hand[j])
			} else if hand[j] == hand[j-1] {
				continue
			} else {
				break
			}

			if len(current) == 2 && hand[len(hand)-1] == OK {
				current = append(current, OK)
			}
			if len(current) > 2 {
				newSeries := make([]int, len(current))
				copy(newSeries, current)
				series = append(series, newSeries)
			}
		}
	}
	return series
}
