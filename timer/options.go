package timer

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type Option func(*State)

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
