package viewport

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/tuiphy/soda"
)

func New(content string, options ...Option) *State {
	state := &State{
		viewport: viewport.New(0, 0),
		content:  content,
		keyMap:   DefaultKeyMap(),
		resizeContent: func(content string, _ soda.Size) string {
			return content
		},
	}

	for _, option := range options {
		option(state)
	}

	state.viewport.KeyMap = viewport.KeyMap(state.keyMap)

	return state
}
