package game

import "sort"

type Player struct {
	hand []int
}

func NewPlayer(tiles []int) *Player {
	return &Player{
		hand: tiles,
	}
}

func (p *Player) AdjustOkeyTiles(okey int) {
	result := make([]int, len(p.hand))
	copy(result, p.hand)

	for i := range result {
		if result[i] == okey {
			result[i] = OK
		}
		if result[i] == 52 {
			result[i] = okey
		}
	}
	sort.Ints(result)
	p.hand = result
}

func (p *Player) GetHand() []int {
	return p.hand
}
