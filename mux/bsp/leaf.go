package bsp

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/cans/mux/window"
	"github.com/tuiphy/soda"
)

var (
	_ Node          = (*Leaf)(nil)
	_ window.Window = (*Leaf)(nil)
)

type LeafOption func(*Leaf)

func WithLeafViewOnly(viewOnly bool) LeafOption {
	return func(leaf *Leaf) {
		leaf.viewOnly = viewOnly
	}
}

func NewLeaf(state soda.State, options ...LeafOption) *Leaf {
	leaf := &Leaf{
		state: state,
	}

	for _, option := range options {
		option(leaf)
	}

	return leaf
}

type Leaf struct {
	state    soda.State
	viewOnly bool

	size soda.Size
}

func (l *Leaf) ViewOnly() bool {
	return l.viewOnly
}

func (l *Leaf) State() soda.State {
	return l.state
}

func (l *Leaf) SetSize(size soda.Size) tea.Cmd {
	l.size = size
	return l.state.SetSize(size)
}

func (l *Leaf) View(layout soda.Layout) string {
	if custom, override := l.state.Layout(); override {
		layout = custom
	}

	return lipgloss.Place(
		l.size.Width,
		l.size.Height,
		layout.Horizontal,
		layout.Vertical,
		l.state.View(layout),
	)
}
