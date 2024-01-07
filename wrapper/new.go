package wrapper

import "github.com/tuiphy/soda"

func New[S soda.State](state S, options ...Option[S]) *State[S] {
	wrapper := &State[S]{
		internal: state,
	}

	for _, option := range options {
		option(wrapper)
	}

	return wrapper
}
