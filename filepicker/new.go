package filepicker

import (
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/tuiphy/soda"
	"os"
)

func New(options ...Option) *State {
	state := &State{
		filePicker: filepicker.New(),
		keyMap:     DefaultKeyMap(),
		onSelect:   soda.Notify,
	}

	wd, err := os.Getwd()
	if err != nil {
		state.filePicker.CurrentDirectory = wd
	}

	for _, option := range options {
		option(state)
	}

	state.filePicker.KeyMap = filepicker.KeyMap(state.keyMap)
	state.filePicker.AutoHeight = false

	return state
}
