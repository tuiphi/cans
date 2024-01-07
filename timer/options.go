package timer

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Option func(*State)

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}

func WithOnTimeout(onTimeout tea.Cmd) Option {
	return func(state *State) {
		state.onTimeout = onTimeout
	}
}

func WithInterval(interval time.Duration) Option {
	return func(state *State) {
		state.interval = interval
	}
}

func WithAutoStart(autostart bool) Option {
	return func(state *State) {
		state.autostart = autostart
	}
}
