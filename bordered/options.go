package bordered

import "github.com/charmbracelet/lipgloss"

type Option func(*State)

func WithBorder(border lipgloss.Border) Option {
	return func(state *State) {
		state.border = border
	}
}

func WithSides(sides ...bool) Option {
	return func(state *State) {
		state.sides = sides
	}
}
