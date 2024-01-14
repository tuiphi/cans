package mux

import (
	"context"
	"math"

	"github.com/tuiphy/cans/mux/window"

	"github.com/charmbracelet/bubbles/key"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	layout    Layout
	activeIdx int
	keyMap    KeyMap
}

func (s *State) windows() []window.Window {
	return s.layout.Windows()
}

func (s *State) viewOnly() bool {
	for _, w := range s.windows() {
		if !w.ViewOnly() {
			return false
		}
	}

	return true
}

func (s *State) activeWindow() window.Window {
	states := s.windows()

	idx := int(math.Abs(float64(s.activeIdx)))
	return states[idx%len(states)]
}

func (s *State) ActivateNext() {
	s.changeActive(1)
	for s.activeWindow().ViewOnly() {
		s.changeActive(1)
	}
}

func (s *State) ActivatePrev() {
	s.changeActive(-1)
	for s.activeWindow().ViewOnly() {
		s.changeActive(-1)
	}
}

func (s *State) changeActive(delta int) {
	if s.viewOnly() {
		return
	}

	nextIdx := s.activeIdx + delta
	if nextIdx < 0 {
		return
	}

	prev := s.activeWindow()

	s.activeIdx = nextIdx % len(s.windows())
	active := s.activeWindow()

	if state, ok := prev.State().(window.ExtendedState); ok {
		state.SetActive(false)
	}

	if state, ok := active.State().(window.ExtendedState); ok {
		state.SetActive(true)
	}
}

// Destroy implements soda.State.
func (s *State) Destroy() {
	for _, window := range s.windows() {
		window.State().Destroy()
	}
}

// Focused implements soda.State.
func (s *State) Focused() bool {
	return s.activeWindow().State().Focused()
}

// Init implements soda.State.
func (s *State) Init(ctx context.Context) tea.Cmd {
	var cmds []tea.Cmd

	for _, w := range s.windows() {
		cmds = append(cmds, w.State().Init(ctx))
	}

	for _, w := range s.windows() {
		if state, ok := w.State().(window.ExtendedState); ok {
			state.SetActive(false)
		}
	}

	if !s.viewOnly() {
		if state, ok := s.activeWindow().State().(window.ExtendedState); ok {
			state.SetActive(true)
		}
	} else {
		s.keyMap.ActivateNext.SetEnabled(false)
		s.keyMap.ActivatePrev.SetEnabled(false)
	}

	return tea.Batch(cmds...)
}

// KeyMap implements soda.State.
func (s *State) KeyMap() help.KeyMap {
	return s.keyMap.with(s.activeWindow().State().KeyMap())
}

// Layout implements soda.State.
func (*State) Layout() (layout soda.Layout, override bool) {
	return
}

// SetSize implements soda.State.
func (s *State) SetSize(size soda.Size) tea.Cmd {
	return s.layout.SetSize(size)
}

// Status implements soda.State.
func (s *State) Status() string {
	return s.activeWindow().State().Status()
}

// Subtitle implements soda.State.
func (s *State) Subtitle() string {
	return s.activeWindow().State().Subtitle()
}

// Title implements soda.State.
func (s *State) Title() string {
	return s.activeWindow().State().Title()
}

// Update implements soda.State.
func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.ActivateNext):
			s.ActivateNext()
			return nil
		case key.Matches(msg, s.keyMap.ActivatePrev):
			s.ActivatePrev()
			return nil
		case !s.viewOnly():
			return s.activeWindow().State().Update(ctx, msg)
		default:
			return nil
		}
	default:
		var cmds []tea.Cmd

		for _, window := range s.windows() {
			cmds = append(cmds, window.State().Update(ctx, msg))
		}

		return tea.Batch(cmds...)
	}
}

// View implements soda.State.
func (s *State) View(layout soda.Layout) string {
	return s.layout.View(layout)
}
