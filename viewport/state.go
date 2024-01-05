package viewport

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	viewport viewport.Model
	content  string

	resizeContent ResizeContentFunc

	keyMap KeyMap
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return false
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	s.viewport.Width = size.Width
	s.viewport.Height = size.Height

	return func() tea.Msg {
		resizedContent := s.resizeContent(s.content, size)
		s.viewport.SetContent(resizedContent)
		return struct{}{}
	}
}

func (s *State) Title() string {
	return "Viewport"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return
}

func (s *State) Status() string {
	return fmt.Sprintf("%3.f%%", s.viewport.ScrollPercent()*100)
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	return s.viewport.Init()
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	s.viewport, cmd = s.viewport.Update(msg)
	return cmd
}

func (s *State) View() string {
	return s.viewport.View()
}
