package mux

func New(layout Layout) *State {
	return &State{
		layout:    layout,
		keyMap:    DefaultKeyMap(),
		activeIdx: 0,
	}
}
