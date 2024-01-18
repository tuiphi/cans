package bsp

import "github.com/tuiphy/soda"

type Direction uint8

const (
	DirectionHorizontal Direction = iota + 1
	DirectionVertical
)

func (d Direction) Split(ratio Ratio, size soda.Size) (soda.Size, soda.Size) {
	switch d {
	case DirectionVertical:
		return size.SplitVertical(float64(ratio))
	case DirectionHorizontal:
		return size.SplitHorizontal(float64(ratio))
	default:
		panic("unreachable")
	}
}
