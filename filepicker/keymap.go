package filepicker

import (
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap(filepicker.DefaultKeyMap())
}

type KeyMap filepicker.KeyMap

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Select,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
