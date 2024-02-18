package bsp

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

type Direction int

const (
	DirectionVertical Direction = iota + 1
	DirectionHorizontal
)

func (d Direction) Join(strs ...string) string {
	switch d {
	case DirectionVertical:
		return lipgloss.JoinVertical(lipgloss.Left, strs...)
	case DirectionHorizontal:
		return lipgloss.JoinHorizontal(lipgloss.Left, strs...)
	default:
		panic("unreachable")
	}
}

func (d Direction) Split(size soda.Size, ratio Ratio) (left, right soda.Size) {
	switch d {
	case DirectionVertical:
		return size.SplitVertical(ratio)
	case DirectionHorizontal:
		return size.SplitHorizontal(ratio)
	default:
		panic("unreachable")
	}
}
