package confirm

import "github.com/tuiphy/soda"

func New(prompt string, options ...Option) *State {
	state := &State{
		onConfirm: soda.Back,
		onCancel:  soda.Back,
		prompt:    prompt,
		keyMap:    KeyMap{},
	}

	for _, option := range options {
		option(state)
	}

	return state
}
