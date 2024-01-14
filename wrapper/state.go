package wrapper

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

type State[S soda.State] struct {
	internal S

	destroy  DestroyFunc
	focused  FocusedFunc
	setSize  SetSizeFunc
	title    TitleFunc
	subtitle SubtitleFunc
	layout   LayoutFunc
	status   StatusFunc
	keyMap   KeyMapFunc
	init     InitFunc
	update   UpdateFunc
	view     ViewFunc
}

func (s *State[S]) Destroy() {
	if s.destroy != nil {
		s.destroy()
	} else {
		s.internal.Destroy()
	}
}

func (s *State[S]) Focused() bool {
	if s.focused != nil {
		return s.focused()
	}
	return s.internal.Focused()
}

func (s *State[S]) SetSize(size soda.Size) tea.Cmd {
	if s.setSize != nil {
		return s.setSize(size)
	}
	return s.internal.SetSize(size)
}

func (s *State[S]) Title() string {
	if s.title != nil {
		return s.title()
	}
	return s.internal.Title()
}

func (s *State[S]) Subtitle() string {
	if s.subtitle != nil {
		return s.subtitle()
	}
	return s.internal.Subtitle()
}

func (s *State[S]) Layout() (layout soda.Layout, override bool) {
	if s.layout != nil {
		return s.layout()
	}
	return s.internal.Layout()
}

func (s *State[S]) Status() string {
	if s.status != nil {
		return s.status()
	}
	return s.internal.Status()
}

func (s *State[S]) KeyMap() help.KeyMap {
	if s.keyMap != nil {
		return s.keyMap()
	}
	return s.internal.KeyMap()
}

func (s *State[S]) Init(ctx context.Context) tea.Cmd {
	if s.init != nil {
		return s.init(ctx)
	}
	return s.internal.Init(ctx)
}

func (s *State[S]) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	if s.update != nil {
		return s.update(ctx, msg)
	}
	return s.internal.Update(ctx, msg)
}

func (s *State[S]) View(layout soda.Layout) string {
	if s.view != nil {
		return s.view(layout)
	}
	return s.internal.View(layout)
}
