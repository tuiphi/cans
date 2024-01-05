package viewport

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap(viewport.DefaultKeyMap())
}

type KeyMap viewport.KeyMap

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{
			k.PageDown,
			k.PageUp,
		},
		{
			k.HalfPageUp,
			k.HalfPageDown,
		},
	}
}
