package bsp

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

type Node interface {
	SetSize(size soda.Size) tea.Cmd
	View(layout soda.Layout) string
}
