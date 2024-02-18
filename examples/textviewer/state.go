package main

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/filepicker"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	filePicker *filepicker.State
}

func (s *State) Destroy() {
	s.filePicker.Destroy()
}

func (s *State) Focused() bool {
	return s.filePicker.Focused()
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	return s.filePicker.SetSize(size)
}

func (s *State) Title() string {
	return "Select a file"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return s.filePicker.Layout()
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.filePicker.KeyMap()
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	return s.filePicker.Init(ctx)
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	return s.filePicker.Update(ctx, msg)
}

func (s *State) View(layout soda.Layout) string {
	return s.filePicker.View(layout)
}
