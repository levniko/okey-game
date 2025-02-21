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
		{0, Yellow, 1, "Yellow-1"},
		{12, Yellow, 13, "Yellow-13"},
		{13, Blue, 1, "Blue-1"},
		{25, Blue, 13, "Blue-13"},
		{26, Black, 1, "Black-1"},
		{38, Black, 13, "Black-13"},
		{39, Red, 1, "Red-1"},
		{51, Red, 13, "Red-13"},
		{52, Fake, 0, "Fake Okey"},
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
		t.Error("IsSameNumber returned false for tiles with the same number but different colors")
	}
	if tile1.IsSameNumber(tile3) {
		t.Error("IsSameNumber returned true for tiles with different numbers")
	}
	if tile1.IsSameNumber(tile4) {
		t.Error("IsSameNumber returned true when compared with a fake okey")
	}
}
