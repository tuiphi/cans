package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SuggestionsSupplier func(value string) []string

type (
	ValidateFunc = textinput.ValidateFunc
	OnSubmitFunc func(value string) tea.Cmd
)

type EchoMode = textinput.EchoMode

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

func WithValidate(f ValidateFunc) Option {
	return func(state *State) {
		state.textInput.Validate = f
	}
}

func WithCharLimit(limit int) Option {
	return func(state *State) {
		state.textInput.CharLimit = limit
	}
}

func WithEchoMode(echoMode EchoMode) Option {
	return func(state *State) {
		state.textInput.EchoMode = echoMode
	}
}

func WithSuggestions(supplier SuggestionsSupplier) Option {
	return func(state *State) {
		state.suggestionsSupplier = supplier

		state.textInput.ShowSuggestions = true
		state.textInput.SetSuggestions(supplier(""))
	}
}
