package timer

import (
	"github.com/charmbracelet/bubbles/timer"
	"time"
)

func New(timeout time.Duration, options ...Option) *State {
	state := &State{
		timeout:   timeout,
		autostart: true,
		interval:  time.Second,
		keyMap:    DefaultKeyMap(),
	}

	for _, option := range options {
		option(state)
	}

	state.timer = timer.NewWithInterval(state.timeout, state.interval)

	return state
}
