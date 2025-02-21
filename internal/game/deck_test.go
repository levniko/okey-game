package game

import (
	"testing"
)

func TestDistributeTiles(t *testing.T) {
	deck := NewDeck()
	players := deck.DistributeTiles()

	if len(players) != 4 {
		t.Errorf("Number of players %d, expected: 4", len(players))
	}

	totalTiles := 0
	fifteenTilePlayerCount := 0

	for _, player := range players {
		tileCount := len(player.Tiles)
		totalTiles += tileCount

		if tileCount == 15 {
			fifteenTilePlayerCount++
		} else if tileCount != 14 {
			t.Errorf("Player %d has %d tiles, expected: 14 or 15",
				player.ID, tileCount)
		}
	}

	if fifteenTilePlayerCount != 1 {
		t.Error("There should be exactly 1 player with 15 tiles")
	}

	if totalTiles != 57 {
		t.Errorf("Total number of distributed tiles %d, expected: 57", totalTiles)
	}
}
