package mux

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/mux/window"
	"github.com/tuiphy/soda"
)

type Layout interface {
	SetSize(size soda.Size) tea.Cmd
	View(layout soda.Layout) string
	Windows() []window.Window
}
