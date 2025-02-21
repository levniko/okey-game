package game

import (
	"testing"
)

func TestGroupScore(t *testing.T) {
	tiles := []Tile{
		{Number: 5, Color: Yellow},
		{Number: 5, Color: Blue},
		{Number: 5, Color: Black},
	}
	okey := Tile{Number: 0, Color: Fake}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 3 {
		t.Errorf("Beklenen puan 3, alınan puan: %d", score)
	}

	foundGroup := false
	for _, comb := range hand.Combinations {
		if comb.Type == Group && len(comb.Tiles) == 3 {
			foundGroup = true
			break
		}
	}
	if !foundGroup {
		t.Error("Grup kombinasyonu bulunamadı")
	}
}

func TestSeriesScore(t *testing.T) {
	tiles := []Tile{
		{Number: 3, Color: Red},
		{Number: 4, Color: Red},
		{Number: 5, Color: Red},
	}
	okey := Tile{Number: 0, Color: Fake}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 3 {
		t.Errorf("Beklenen puan 3, alınan puan: %d", score)
	}

	foundSeries := false
	for _, comb := range hand.Combinations {
		if comb.Type == Series && len(comb.Tiles) == 3 {
			foundSeries = true
			break
		}
	}
	if !foundSeries {
		t.Error("Seri kombinasyonu bulunamadı")
	}
}

func TestPairScore(t *testing.T) {
	tiles := []Tile{
		{Number: 7, Color: Blue},
		{Number: 7, Color: Blue},
	}
	okey := Tile{Number: 0, Color: Fake}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 2 {
		t.Errorf("Beklenen puan 2, alınan puan: %d", score)
	}

	foundPair := false
	for _, comb := range hand.Combinations {
		if comb.Type == Pair && len(comb.Tiles) == 2 {
			foundPair = true
			break
		}
	}
	if !foundPair {
		t.Error("Çift kombinasyonu bulunamadı")
	}
}

func TestOkeyReplacement(t *testing.T) {
	tiles := []Tile{
		{Number: 0, Color: Fake},
		{Number: 5, Color: Yellow},
	}
	okey := Tile{Number: 3, Color: Red}
	hand := NewHand(tiles, okey)

	if hand.Tiles[0] != okey {
		t.Error("Sahte okey gerçek okeye dönüştürülmedi")
	}
}

func TestMixedCombinations(t *testing.T) {
	tiles := []Tile{
		{Number: 5, Color: Yellow},
		{Number: 5, Color: Blue},
		{Number: 5, Color: Black},
		{Number: 3, Color: Red},
		{Number: 4, Color: Red},
		{Number: 5, Color: Red},
	}
	okey := Tile{Number: 0, Color: Fake}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 6 {
		t.Errorf("Beklenen puan 6, alınan puan: %d", score)
	}
}

func TestEdgeCases(t *testing.T) {
	tiles := []Tile{
		{Number: 7, Color: Yellow},
		{Number: 7, Color: Blue},
		{Number: 7, Color: Black},
		{Number: 7, Color: Red},
	}
	okey := Tile{Number: 0, Color: Fake}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 4 {
		t.Errorf("Beklenen puan 4, alınan puan: %d", score)
	}
}

func TestOkeyJoker(t *testing.T) {
	tiles := []Tile{
		{Number: 3, Color: Red},
		{Number: 5, Color: Red},
		{IsOkey: true},
	}
	okey := Tile{Number: 4, Color: Red}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()

	if score != 3 {
		t.Errorf("Okey joker testi başarısız. Beklenen: 3, Alınan: %d", score)
	}
}

func TestOkeyAsJokerInGroup(t *testing.T) {
	tiles := []Tile{
		{Number: 1, Color: Blue},
		{Number: 1, Color: Red},
		{IsOkey: true},
	}
	okey := Tile{Number: 1, Color: Yellow}
	hand := NewHand(tiles, okey)
	score := hand.CalculateScore()
	if score != 3 {
		t.Errorf("Beklenen skor 3, alınan: %d", score)
	}
}
