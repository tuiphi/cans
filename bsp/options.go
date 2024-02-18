package bsp

type Option func(*State)

func WithBranch(branch Branch) Option {
	return func(state *State) {
		state.branch = branch
	}
}

func WithCycle(cycle bool) Option {
	return func(state *State) {
		state.cycle = cycle
	}
}
