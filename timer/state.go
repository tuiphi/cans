package timer

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
	"time"
)

var _ soda.State = (*State)(nil)

type State struct {
	interval time.Duration
	timeout  time.Duration

	autostart bool

	timer     timer.Model
	onTimeout tea.Cmd

	keyMap KeyMap
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return false
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	return nil
}

func (s *State) Title() string {
	return "Timer"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return soda.Layout{
		Horizontal: lipgloss.Center,
		Vertical:   lipgloss.Center,
	}, true
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	s.keyMap.Start.SetEnabled(!s.autostart)
	s.keyMap.Stop.SetEnabled(s.autostart)

	cmds := []tea.Cmd{s.timer.Init()}
	if !s.autostart {
		cmds = append(cmds, s.timer.Stop())
	}

	return tea.Batch(cmds...)
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		s.timer, cmd = s.timer.Update(msg)
		return cmd
	case timer.StartStopMsg:
		var cmd tea.Cmd
		s.timer, cmd = s.timer.Update(msg)
		s.keyMap.Stop.SetEnabled(s.timer.Running())
		s.keyMap.Start.SetEnabled(!s.timer.Running())
		return cmd
	case timer.TimeoutMsg:
		s.keyMap.Start.SetEnabled(false)
		s.keyMap.Stop.SetEnabled(false)
		return s.onTimeout
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Start, s.keyMap.Stop):
			return s.timer.Toggle()
		case key.Matches(msg, s.keyMap.Reset):
			s.timer.Timeout = s.timeout
			return s.timer.Stop()
		}
	}

	return nil
}

func (s *State) View() string {
	return s.timer.View()
}
