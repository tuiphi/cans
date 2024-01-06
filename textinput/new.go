package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func New(options ...Option) *State {
	state := &State{
		textInput: textinput.New(),
		keyMap:    DefaultKeyMap(),
		onSubmit:  nil,
	}

	for _, option := range options {
		option(state)
	}

	return state
}
