package confirm

import tea "github.com/charmbracelet/bubbletea"

type Option func(*State)

func WithOnConfirm(onConfirm tea.Cmd) Option {
	return func(state *State) {
		state.onConfirm = onConfirm
	}
}

func WithOnCancel(onCancel tea.Cmd) Option {
	return func(state *State) {
		state.onCancel = onCancel
	}
}

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}

func WithPrompt(prompt string) Option {
	return func(state *State) {
		state.prompt = prompt
	}
}
