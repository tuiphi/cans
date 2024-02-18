package bsp

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	branch Branch

	activeIndex int

	keyMap KeyMap
	size   soda.Size
}

func (s *State) Active() (soda.State, bool) {
	states := s.branch.States()
	if s.activeIndex < 0 || s.activeIndex >= len(states) {
		return nil, false
	}

	return states[s.activeIndex], true
}

func (s *State) ActivateNext() tea.Cmd {
	return s.activateDelta(1)
}

func (s *State) ActivatePrev() tea.Cmd {
	return s.activateDelta(-1)
}

func (s *State) activateDelta(delta int) tea.Cmd {
	index := s.clampIndex(s.activeIndex + delta)
	return s.activate(index)
}

func (s *State) activate(index int) tea.Cmd {
	var cmds []tea.Cmd

	handle := func() {
		if active, ok := s.Active(); ok {
			if activater, ok := active.(Activater); ok {
				cmds = append(cmds, activater.Deactivate())
			}
		}
	}

	handle()
	s.activeIndex = index
	handle()

	return tea.Sequence(cmds...)
}

func (s *State) clampIndex(index int) int {
	if index < 0 {
		return len(s.branch.States()) - 1
	}

	return index % len(s.branch.States())
}

// Destroy implements soda.State.
func (s *State) Destroy() {
	s.branch.Destroy()
}

// Focused implements soda.State.
func (s *State) Focused() bool {
	if active, ok := s.Active(); ok {
		return active.Focused()
	}

	return false
}

// Init implements soda.State.
func (s *State) Init(ctx context.Context) tea.Cmd {
	return s.branch.Init(ctx)
}

// KeyMap implements soda.State.
func (s *State) KeyMap() help.KeyMap {
	if active, ok := s.Active(); ok {
		return s.keyMap.with(active.KeyMap())
	}

	return s.keyMap
}

// Layout implements soda.State.
func (*State) Layout() (layout soda.Layout, override bool) {
	return
}

// SetSize implements soda.State.
func (s *State) SetSize(size soda.Size) tea.Cmd {
	s.size = size
	return s.branch.SetSize(size)
}

// Status implements soda.State.
func (s *State) Status() string {
	if active, ok := s.Active(); ok {
		return active.Status()
	}

	return ""
}

// Subtitle implements soda.State.
func (s *State) Subtitle() string {
	if active, ok := s.Active(); ok {
		return active.Subtitle()
	}

	return ""
}

// Title implements soda.State.
func (s *State) Title() string {
	if active, ok := s.Active(); ok {
		return active.Title()
	}

	return ""
}

// Update implements soda.State.
func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case BranchSetMsg:
		s.branch = msg.Branch
		s.activeIndex = s.clampIndex(s.activeIndex)
		return nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Next):
			return s.ActivateNext()
		case key.Matches(msg, s.keyMap.Prev):
			return s.ActivatePrev()
		default:
			if active, ok := s.Active(); ok {
				return active.Update(ctx, msg)
			}

			return nil
		}
	default:
		return s.branch.Update(ctx, msg)
	}
}

// View implements soda.State.
func (s *State) View(layout soda.Layout) string {
	return s.branch.View(layout)
}
