package main

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*SizeHolder)(nil)

type SizeHolder struct {
	title string
	size  soda.Size
}

func (s *SizeHolder) Destroy() {
}

func (s *SizeHolder) Focused() bool {
	return false
}

func (s *SizeHolder) SetSize(size soda.Size) tea.Cmd {
	s.size = size
	return nil
}

func (s *SizeHolder) Title() string {
	return s.title
}

func (s *SizeHolder) Subtitle() string {
	return ""
}

func (s *SizeHolder) Layout() (layout soda.Layout, override bool) {
	return soda.Layout{
		Horizontal: lipgloss.Center,
		Vertical:   lipgloss.Center,
	}, true
}

func (s *SizeHolder) Status() string {
	return ""
}

func (s *SizeHolder) KeyMap() help.KeyMap {
	return help.KeyMap(nil)
}

func (s *SizeHolder) Init(ctx context.Context) tea.Cmd {
	return nil
}

func (s *SizeHolder) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	return nil
}

func (s *SizeHolder) View(soda.Layout) string {
	return s.size.String()
}
