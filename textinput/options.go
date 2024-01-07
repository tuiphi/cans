package textinput

import tea "github.com/charmbracelet/bubbletea"

type OnSubmitFunc func(value string) tea.Cmd

type SuggestionsSupplier func(value string) []string

type Option func(*State)

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}

func WithOnSubmit(submitFunc OnSubmitFunc) Option {
	return func(state *State) {
		state.onSubmit = submitFunc
	}
}

func WithPlaceholder(placeholder string) Option {
	return func(state *State) {
		state.textInput.Placeholder = placeholder
	}
}

func WithSuggestions(supplier SuggestionsSupplier) Option {
	return func(state *State) {
		state.suggestionsSupplier = supplier

		state.textInput.ShowSuggestions = true
		state.textInput.SetSuggestions(supplier(""))
	}
}
