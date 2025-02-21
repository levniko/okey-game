package game

import (
	"testing"
)

func TestNewTile(t *testing.T) {
	tests := []struct {
		id          int
		wantColor   Color
		wantNumber  int
		description string
	}{
		{0, Yellow, 1, "Sarı-1"},
		{12, Yellow, 13, "Sarı-13"},
		{13, Blue, 1, "Mavi-1"},
		{25, Blue, 13, "Mavi-13"},
		{26, Black, 1, "Siyah-1"},
		{38, Black, 13, "Siyah-13"},
		{39, Red, 1, "Kırmızı-1"},
		{51, Red, 13, "Kırmızı-13"},
		{52, Fake, 0, "Sahte Okey"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tile := NewTile(tt.id)
			if tile.Color != tt.wantColor {
				t.Errorf("NewTile(%d).Color = %v, want %v", tt.id, tile.Color, tt.wantColor)
			}
			if tile.Number != tt.wantNumber {
				t.Errorf("NewTile(%d).Number = %v, want %v", tt.id, tile.Number, tt.wantNumber)
			}
		})
	}
}

func TestIsSameNumber(t *testing.T) {
	tile1 := NewTile(0)
	tile2 := NewTile(13)
	tile3 := NewTile(1)
	tile4 := NewTile(52)

	if !tile1.IsSameNumber(tile2) {
		t.Error("Aynı sayılı farklı renkli taşlar için IsSameNumber false döndü")
	}
	if tile1.IsSameNumber(tile3) {
		t.Error("Farklı sayılı taşlar için IsSameNumber true döndü")
	}
	if tile1.IsSameNumber(tile4) {
		t.Error("Sahte okey ile karşılaştırmada IsSameNumber true döndü")
	}
}
