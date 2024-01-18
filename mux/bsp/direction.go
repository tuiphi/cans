package bsp

import "github.com/tuiphy/soda"

type Direction uint8

const (
	DirectionHorizontal Direction = iota + 1
	DirectionVertical
)

func (d Direction) Split(ratio float64, size soda.Size) (soda.Size, soda.Size) {
	switch d {
	case DirectionVertical:
		return size.SplitVertical(ratio)
	case DirectionHorizontal:
		return size.SplitHorizontal(ratio)
	default:
		panic("unreachable")
	}
}
