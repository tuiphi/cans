package viewport

import (
	"github.com/tuiphy/soda"
)

type ResizeContentFunc func(content string, size soda.Size) string

type Option func(*State)

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}

func WithResizeContent(f ResizeContentFunc) Option {
	return func(state *State) {
		state.resizeContent = f
	}
}
