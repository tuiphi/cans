package confirm

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Confirm: key.NewBinding(
			key.WithKeys("enter", "y", "Y"),
			key.WithHelp("enter", "confirm"),
		),
		Cancel: key.NewBinding(
			key.WithKeys("n", "N"),
			key.WithHelp("n", "cancel"),
		),
	}
}

type KeyMap struct {
	Confirm, Cancel key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Confirm,
		k.Cancel,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
