package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	l := list.DefaultKeyMap()
	l.CancelWhileFiltering = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	)

	return KeyMap{
		List: l,
		Select: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "select"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit"),
		),
	}
}

type KeyMap struct {
	List list.KeyMap
	Select,
	Submit key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Select,
		k.Submit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
