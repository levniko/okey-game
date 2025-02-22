package game

import (
	"math/rand"
	"time"
)

type Deck struct {
	tiles     []int
	indicator int
	okey      int
	players   []*Player
}

func NewDeck() *Deck {
	rand.Seed(time.Now().UnixNano())
	d := &Deck{
		players: make([]*Player, 4),
	}
	d.tiles = createAndShuffleTiles()
	d.indicator = selectIndicator(d.tiles)
	d.okey = findOkey(d.indicator)
	d.distributeTilesToPlayers()
	return d
}

func (d *Deck) distributeTilesToPlayers() {
	playerWith15 := rand.Intn(4)

	tileStart := 0
	for i := 0; i < 4; i++ {
		var tileCount int
		if i == playerWith15 {
			tileCount = 15
		} else {
			tileCount = 14
		}

		d.players[i] = NewPlayer(d.tiles[tileStart : tileStart+tileCount])
		d.players[i].AdjustOkeyTiles(d.okey)
		tileStart += tileCount
	}
}

func (d *Deck) GetPlayers() []*Player {
	return d.players
}

func (d *Deck) GetPlayerWith15() int {
	for i, player := range d.players {
		if len(player.GetHand()) == 15 {
			return i
		}
	}
	return -1
}

func (d *Deck) GetTiles() []int {
	return d.tiles
}

func (d *Deck) GetIndicator() int {
	return d.indicator
}

func (d *Deck) GetOkey() int {
	return d.okey
}

func createAndShuffleTiles() []int {
	tiles := make([]int, TotalTiles)
	for i := 0; i < TotalTiles; i++ {
		tiles[i] = i % DeckSize
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})
	return tiles
}

func selectIndicator(tiles []int) int {
	if tiles[TotalTiles-1] != 52 {
		return tiles[TotalTiles-1]
	}
	return tiles[TotalTiles-2]
}

func findOkey(indicator int) int {
	if (indicator+1)%13 == 0 {
		return indicator - 12
	}
	return indicator + 1
}
