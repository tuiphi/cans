package filepicker

type Option func(*State)

func WithDir(dir string) Option {
	return func(state *State) {
		state.filePicker.CurrentDirectory = dir
	}
}

func WithOnSelect(onSelect OnSelectFunc) Option {
	return func(state *State) {
		state.onSelect = onSelect
	}
}

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}
