package game

import "fmt"

type Color int

const (
	Yellow Color = iota
	Blue
	Black
	Red
	Fake
)

func (c Color) String() string {
	switch c {
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Black:
		return "Black"
	case Red:
		return "Red"
	case Fake:
		return "Fake"
	default:
		return "Unknown"
	}
}

type Tile struct {
	ID     int
	Color  Color
	Number int
	IsOkey bool
}

func NewTile(id int) Tile {
	if id == 52 {
		return Tile{
			ID:     id,
			Color:  Fake,
			Number: 0,
			IsOkey: false,
		}
	}

	color := Color(id / 13)
	number := (id % 13) + 1

	return Tile{
		ID:     id,
		Color:  color,
		Number: number,
		IsOkey: false,
	}
}

func (t Tile) String() string {
	if t.Color == Fake {
		return "Fake Okey"
	}
	return fmt.Sprintf("%s-%d", t.Color, t.Number)
}

func (t Tile) IsSameNumber(other Tile) bool {
	return t.Number == other.Number && t.Color != Fake && other.Color != Fake
}

func (t Tile) IsNextNumber(other Tile) bool {
	return t.Color == other.Color && t.Number == other.Number+1
}
