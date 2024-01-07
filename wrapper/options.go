package wrapper

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

type (
	DestroyFunc  func()
	FocusedFunc  func() bool
	SetSizeFunc  func(size soda.Size) tea.Cmd
	TitleFunc    func() string
	SubtitleFunc func() string
	LayoutFunc   func() (layout soda.Layout, override bool)
	StatusFunc   func() string
	KeyMapFunc   func() help.KeyMap
	InitFunc     func(ctx context.Context) tea.Cmd
	UpdateFunc   func(ctx context.Context, msg tea.Msg) tea.Cmd
	ViewFunc     func() string
)

type Option[S soda.State] func(*State[S])

func WithDestroy[S soda.State](destroyFunc DestroyFunc) Option[S] {
	return func(state *State[S]) {
		state.destroy = destroyFunc
	}
}

func WithFocused[S soda.State](focusedFunc FocusedFunc) Option[S] {
	return func(state *State[S]) {
		state.focused = focusedFunc
	}
}

func WithSetSize[S soda.State](setSizeFunc SetSizeFunc) Option[S] {
	return func(state *State[S]) {
		state.setSize = setSizeFunc
	}
}

func WithTitle[S soda.State](titleFunc TitleFunc) Option[S] {
	return func(state *State[S]) {
		state.title = titleFunc
	}
}

func WithSubtitle[S soda.State](subtitleFunc SubtitleFunc) Option[S] {
	return func(state *State[S]) {
		state.subtitle = subtitleFunc
	}
}

func WithLayout[S soda.State](layoutFunc LayoutFunc) Option[S] {
	return func(state *State[S]) {
		state.layout = layoutFunc
	}
}

func WithStatus[S soda.State](statusFunc StatusFunc) Option[S] {
	return func(state *State[S]) {
		state.status = statusFunc
	}
}

func WithKeyMap[S soda.State](keyMapFunc KeyMapFunc) Option[S] {
	return func(state *State[S]) {
		state.keyMap = keyMapFunc
	}
}

func WithInit[S soda.State](initFunc InitFunc) Option[S] {
	return func(state *State[S]) {
		state.init = initFunc
	}
}

func WithUpdate[S soda.State](updateFunc UpdateFunc) Option[S] {
	return func(state *State[S]) {
		state.update = updateFunc
	}
}

func WithView[S soda.State](viewFunc ViewFunc) Option[S] {
	return func(state *State[S]) {
		state.view = viewFunc
	}
}
