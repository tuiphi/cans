package confirm

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	onConfirm, onCancel tea.Cmd

	prompt string

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
	return "Confirm"
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
	return nil
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Confirm):
			return s.onConfirm
		case key.Matches(msg, s.keyMap.Cancel):
			return s.onCancel
		}
	}

	return nil
}

func (s *State) View(soda.Layout) string {
	return s.prompt
}
