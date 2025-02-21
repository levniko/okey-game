package game

import (
	"math/rand"
	"time"
)

type Deck struct {
	tiles     []Tile
	indicator Tile
	okeyTile  Tile
}

func NewDeck() *Deck {
	rand.Seed(time.Now().UnixNano())

	tiles := make([]Tile, 0, 106)
	for i := range 53 {
		tiles = append(tiles, NewTile(i))
		tiles = append(tiles, NewTile(i))
	}

	deck := &Deck{
		tiles: tiles,
	}

	deck.Shuffle()

	deck.setIndicatorAndOkey()

	return deck
}

func (d *Deck) Shuffle() {
	n := len(d.tiles)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.tiles[i], d.tiles[j] = d.tiles[j], d.tiles[i]
	}
}

func (d *Deck) setIndicatorAndOkey() {
	d.indicator = d.tiles[len(d.tiles)-1]
	if d.indicator.ID == 52 {
		d.indicator = d.tiles[len(d.tiles)-2]
	}
	d.okeyTile.Number = (d.indicator.Number % 13) + 1
	d.okeyTile.Color = d.indicator.Color
}

func (d *Deck) DistributeTiles() []Player {
	players := make([]Player, 4)
	for i := range players {
		players[i] = Player{ID: i}
	}

	luckyPlayer := rand.Intn(4)
	currentTile := 0

	for i := range players {
		tileCount := 14
		if i == luckyPlayer {
			tileCount = 15
		}

		players[i].Tiles = make([]Tile, tileCount)
		for j := 0; j < tileCount; j++ {
			players[i].Tiles[j] = d.tiles[currentTile]
			currentTile++
		}
	}

	return players
}

func (d *Deck) GetIndicator() Tile {
	return d.indicator
}

func (d *Deck) GetOkeyTile() Tile {
	return d.okeyTile
}
