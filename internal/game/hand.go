package game

import (
	"math"
	"slices"
)

func CalculateHandScore(hand []int) int {
	series := findAllSeries(hand)
	groups := findAllGroups(hand)
	series = append(series, groups...)

	maxScore := 0
	for _, combination := range series {
		remaining := make([]int, 0)
		handMap := make(map[int]bool)
		for _, tile := range hand {
			handMap[tile] = true
		}

		for _, tile := range combination {
			delete(handMap, tile)
		}

		for tile := range handMap {
			remaining = append(remaining, tile)
		}

		if len(remaining) == 0 {
			return len(combination)
		}

		score := CalculateHandScore(remaining)
		maxScore = int(math.Max(float64(score+len(combination)), float64(maxScore)))
	}

	return maxScore
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
	groups := make([][]int, 13)
	for i := range groups {
		groups[i] = make([]int, 0)
		if hand[len(hand)-1] == OK {
			groups[i] = append(groups[i], OK)
		}
		if hand[len(hand)-2] == OK {
			groups[i] = append(groups[i], OK)
		}
	}

	for _, tile := range hand {
		num := tile % 13
		exists := false
		for _, t := range groups[num] {
			if t == tile {
				exists = true
				break
			}
		}
		if !exists {
			groups[num] = append(groups[num], tile)
		}
	}

	var validGroups [][]int
	for _, group := range groups {
		for len(group) > 4 {
			for i, tile := range group {
				if tile == OK {
					group = slices.Delete(group, i, i+1)
					break
				}
			}
		}

		if len(group) >= 3 {
			validGroups = append(validGroups, slices.Clone(group))
		}
		if len(group) == 4 {
			for i := range group {
				newGroup := make([]int, 0)
				for j, tile := range group {
					if i != j {
						newGroup = append(newGroup, tile)
					}
				}
				validGroups = append(validGroups, newGroup)
			}
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
