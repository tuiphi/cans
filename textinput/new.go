package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func New(options ...Option) *State {
	state := &State{
		textInput: textinput.New(),
		keyMap:    DefaultKeyMap(),
		onSubmit: func(value string) tea.Cmd {
			return nil
		},
		suggestionsSupplier: func(string) []string {
			return nil
		},
	}

	for _, option := range options {
		option(state)
	}

	state.textInput.KeyMap = state.keyMap.TextInput

	return state
}
