package model

type Color int

const (
	COLORLESS Color = iota
	WHITE
	BLUE
	BLACK
	RED
	GREEN
	MULTI
)

type Cost string

type ManaCost map[Color]Cost

type Card struct {
	Name     string
	Color    Color
	ManaCost ManaCost
}
