package textinput

import tea "github.com/charmbracelet/bubbletea"

type OnSubmitFunc func(value string) tea.Cmd

type Option func(*State)

func WithOnSubmit(submitFunc OnSubmitFunc) Option {
	return func(state *State) {
		state.onSubmit = submitFunc
	}
}
