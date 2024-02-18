package nop

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	keyMap KeyMap
}

// Destroy implements soda.State.
func (State) Destroy() {
}

// Focused implements soda.State.
func (State) Focused() bool {
	return false
}

// Init implements soda.State.
func (State) Init(context.Context) tea.Cmd {
	return nil
}

// KeyMap implements soda.State.
func (s State) KeyMap() help.KeyMap {
	return s.keyMap
}

// Layout implements soda.State.
func (State) Layout() (layout soda.Layout, override bool) {
	return
}

// SetSize implements soda.State.
func (State) SetSize(soda.Size) tea.Cmd {
	return nil
}

// Status implements soda.State.
func (State) Status() string {
	return ""
}

// Subtitle implements soda.State.
func (State) Subtitle() string {
	return ""
}

// Title implements soda.State.
func (State) Title() string {
	return ""
}

// Update implements soda.State.
func (State) Update(context.Context, tea.Msg) tea.Cmd {
	return nil
}

// View implements soda.State.
func (State) View(soda.Layout) string {
	return ""
}
