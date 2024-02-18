package nop

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

type KeyMap struct{}

// FullHelp implements help.KeyMap.
func (KeyMap) FullHelp() [][]key.Binding {
	return nil
}

// ShortHelp implements help.KeyMap.
func (KeyMap) ShortHelp() []key.Binding {
	return nil
}
