package filepicker

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type OnSelectFunc func(path string) tea.Cmd

type State struct {
	filePicker filepicker.Model
	keyMap     KeyMap

	onSelect OnSelectFunc
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return false
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	// NOTE: this is correct
	s.filePicker.Height = size.Height - 1
	return nil
}

func (s *State) Title() string {
	return "File picker"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	return s.filePicker.Init()
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	s.filePicker, cmd = s.filePicker.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Select):
			if didSelectDisabled, path := s.filePicker.DidSelectDisabledFile(msg); didSelectDisabled {
				return soda.SendError(fmt.Errorf("%s is not valid", path))
			}

			if didSelect, path := s.filePicker.DidSelectFile(msg); didSelect {
				return s.onSelect(path)
			}
		}
	}

	return cmd
}

func (s *State) View() string {
	return s.filePicker.View()
}
