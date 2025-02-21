package game

import (
	"sort"
)

type CombinationType int

const (
	Series CombinationType = iota
	Group
	Pair
)

func (c CombinationType) String() string {
	switch c {
	case Series:
		return "Series"
	case Group:
		return "Group"
	case Pair:
		return "Pair"
	default:
		return "Unknown"
	}
}

type Combination struct {
	Type  CombinationType
	Tiles []Tile
}

type Hand struct {
	Tiles        []Tile
	Okey         Tile
	Combinations []Combination
	MaxScore     int
}

func NewHand(tiles []Tile, okey Tile) *Hand {
	h := &Hand{
		Tiles: make([]Tile, len(tiles)),
		Okey:  okey,
	}
	copy(h.Tiles, tiles)
	h.replaceOkeys()
	h.sortTiles()
	h.CalculateScore()
	return h
}

func (h *Hand) replaceOkeys() {
    for i := range h.Tiles {
        if h.Tiles[i].ID == 52 || h.Tiles[i].IsOkey {
            h.Tiles[i] = h.Okey
        }
    }
}

func (h *Hand) sortTiles() {
	sort.Slice(h.Tiles, func(i, j int) bool {
		if h.Tiles[i].Color == h.Tiles[j].Color {
			return h.Tiles[i].Number < h.Tiles[j].Number
		}
		return h.Tiles[i].Color < h.Tiles[j].Color
	})
}

func (h *Hand) CalculateScore() int {
	h.MaxScore = h.calculateOptimalScore(h.Tiles)
	_, pairCombinations := h.calculatePairScore()
	h.Combinations = append(h.Combinations, pairCombinations...)
	return h.MaxScore
}

func (h *Hand) calculateOptimalScore(tiles []Tile) int {
	allCombos := h.findAllCombinations(tiles)
	maxScore := 0

	for _, comboSet := range generateComboSets(allCombos) {
		if isValidComboSet(comboSet) {
			score := calculateComboSetScore(comboSet)
			if score > maxScore {
				maxScore = score
				h.Combinations = comboSet
			}
		}
	}

	pairScore, pairCombinations := h.calculatePairScore()
	h.Combinations = append(h.Combinations, pairCombinations...)
	if pairScore > maxScore {
		maxScore = pairScore
	}

	return maxScore
}

func (h *Hand) calculatePairScore() (int, []Combination) {
	h.sortTiles()
	pairs := 0
	var pairCombos []Combination

	for i := 0; i < len(h.Tiles)-1; i++ {
		if h.Tiles[i].Number == h.Tiles[i+1].Number && h.Tiles[i].Color == h.Tiles[i+1].Color {
			pairs++
			pairCombos = append(pairCombos, Combination{
				Type:  Pair,
				Tiles: []Tile{h.Tiles[i], h.Tiles[i+1]},
			})
			i++ 
		}
	}

	return pairs * 2, pairCombos
}

func generateComboSets(combos []Combination) [][]Combination {
	var sets [][]Combination
	sets = append(sets, combos)
	return sets
}

func isValidComboSet(comboSet []Combination) bool {
	used := make(map[int]bool)
	for _, combo := range comboSet {
		for _, tile := range combo.Tiles {
			if used[tile.ID] {
				return false
			}
			used[tile.ID] = true
		}
	}
	return true
}

func calculateComboSetScore(comboSet []Combination) int {
    score := 0
    for _, combo := range comboSet {
        score += len(combo.Tiles)
    }
    return score
}

func (h *Hand) findAllCombinations(tiles []Tile) []Combination {
	var combos []Combination
	combos = append(combos, h.findGroupsWithOkey(tiles)...)
	combos = append(combos, h.findSeriesWithOkey(tiles)...)
	return combos
}

func (h *Hand) findGroupsWithOkey(tiles []Tile) []Combination {
	groups := make(map[int][]Tile)
	for _, t := range tiles {
		num := t.Number
		if t.IsOkey {
			num = h.Okey.Number 
		}
		groups[num] = append(groups[num], t)
	}

	var combos []Combination
	for _, group := range groups {
		if len(group) >= 3 {
			colorMap := make(map[Color]bool)
			validGroup := []Tile{}
			for _, tile := range group {
				if !colorMap[tile.Color] {
					colorMap[tile.Color] = true
					validGroup = append(validGroup, tile)
				}
			}
			if len(validGroup) >= 3 {
				combos = append(combos, Combination{
					Type:  Group,
					Tiles: validGroup[:min(len(validGroup), 4)],
				})
			}
		}
	}
	return combos
}

func (h *Hand) findSeriesWithOkey(tiles []Tile) []Combination {
	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i].Color == tiles[j].Color {
			return tiles[i].Number < tiles[j].Number
		}
		return tiles[i].Color < tiles[j].Color
	})

	var combos []Combination
	currentSeries := []Tile{tiles[0]}
	okeyUsed := false

	for i := 1; i < len(tiles); i++ {
		prev := currentSeries[len(currentSeries)-1]
		curr := tiles[i]

		if curr.Color == prev.Color && curr.Number == prev.Number+1 {
			currentSeries = append(currentSeries, curr)
		} else if h.hasOkey() && !okeyUsed && curr.Color == prev.Color && curr.Number == prev.Number+2 {
			currentSeries = append(currentSeries, h.Okey, curr)
			okeyUsed = true
		} else {
			if len(currentSeries) >= 3 {
				combos = append(combos, Combination{
					Type:  Series,
					Tiles: currentSeries,
				})
			}
			currentSeries = []Tile{curr}
			okeyUsed = false
		}
	}

	if len(currentSeries) >= 3 {
		combos = append(combos, Combination{
			Type:  Series,
			Tiles: currentSeries,
		})
	}
	return combos
}

func (h *Hand) hasOkey() bool {
	for _, t := range h.Tiles {
		if t.IsOkey || t.Color == Fake {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
