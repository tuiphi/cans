package bsp

import "github.com/tuiphy/soda"

type Direction uint8

const (
	DirectionHorizontal Direction = iota + 1
	DirectionVertical
)

func (d Direction) Split(size soda.Size) (soda.Size, soda.Size) {
	const ratio = 0.5

	switch d {
	case DirectionVertical:
		return size.SplitVertical(ratio)
	case DirectionHorizontal:
		return size.SplitHorizontal(ratio)
	default:
		panic("unreachable")
	}
}
