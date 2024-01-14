package bordered

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	size                   soda.Size
	inner                  soda.State
	sides                  []bool
	border                 lipgloss.Border
	foreground, background lipgloss.Color
}

// Destroy implements soda.State.
func (s *State) Destroy() {
	s.inner.Destroy()
}

// Focused implements soda.State.
func (s *State) Focused() bool {
	return s.inner.Focused()
}

// Init implements soda.State.
func (s *State) Init(ctx context.Context) tea.Cmd {
	s.SetActive(false)
	return s.inner.Init(ctx)
}

// KeyMap implements soda.State.
func (s *State) KeyMap() help.KeyMap {
	return s.inner.KeyMap()
}

// Layout implements soda.State.
func (s *State) Layout() (layout soda.Layout, override bool) {
	return s.inner.Layout()
}

// SetSize implements soda.State.
func (s *State) SetSize(size soda.Size) tea.Cmd {
	size.Width -= s.border.GetLeftSize() + s.border.GetRightSize()
	size.Height -= s.border.GetBottomSize() + s.border.GetTopSize()

	s.size = size

	return s.inner.SetSize(size)
}

// Status implements soda.State.
func (s *State) Status() string {
	return s.inner.Status()
}

// Subtitle implements soda.State.
func (s *State) Subtitle() string {
	return s.inner.Subtitle()
}

// Title implements soda.State.
func (s *State) Title() string {
	return s.inner.Title()
}

// Update implements soda.State.
func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	return s.inner.Update(ctx, msg)
}

func (s *State) SetForeground(color lipgloss.Color) {
	s.foreground = color
}

func (s *State) SetBackground(color lipgloss.Color) {
	s.background = color
}

func (s *State) SetActive(active bool) {
	if active {
		s.SetForeground("#DCDCDC")
	} else {
		s.SetForeground("#A9A9A9")
	}
}

// View implements soda.State.
func (s *State) View(layout soda.Layout) string {
	border := lipgloss.
		NewStyle().
		Border(s.border, s.sides...).
		BorderForeground(s.foreground).
		BorderBackground(s.background)

	view := lipgloss.Place(
		s.size.Width,
		s.size.Height,
		layout.Horizontal,
		layout.Vertical,
		s.inner.View(layout),
	)

	return border.Render(view)
}
