package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func New(options ...Option) *State {
	state := &State{
		textInput: textinput.New(),
		keyMap:    DefaultKeyMap(),
		onSubmit:  nil,
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
