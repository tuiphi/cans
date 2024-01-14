package bordered

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

func New(state soda.State, options ...Option) *State {
	s := &State{
		inner:      state,
		border:     lipgloss.NormalBorder(),
		foreground: lipgloss.Color(""),
		background: lipgloss.Color(""),
		sides:      []bool{true, true, true, true},
	}

	for _, option := range options {
		option(s)
	}

	return s
}
