package game

import (
	"testing"
)

func TestDistributeTiles(t *testing.T) {
	deck := NewDeck()
	players := deck.DistributeTiles()

	if len(players) != 4 {
		t.Errorf("Oyuncu sayısı %d, olması gereken: 4", len(players))
	}

	totalTiles := 0
	fifteenTilePlayerCount := 0

	for _, player := range players {
		tileCount := len(player.Tiles)
		totalTiles += tileCount

		if tileCount == 15 {
			fifteenTilePlayerCount++
		} else if tileCount != 14 {
			t.Errorf("Oyuncu %d'nin taş sayısı %d, olması gereken: 14 veya 15",
				player.ID, tileCount)
		}
	}

	if fifteenTilePlayerCount != 1 {
		t.Error("15 taşı olan oyuncu sayısı 1 olmalı")
	}

	if totalTiles != 57 {
		t.Errorf("Toplam dağıtılan taş sayısı %d, olması gereken: 57", totalTiles)
	}
}
