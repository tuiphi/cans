package textinput

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		TextInput: textinput.DefaultKeyMap,
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "submit"),
		),
		Focus: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "focus"),
		),
		Blur: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "blur"),
		),
	}
}

type KeyMap struct {
	TextInput textinput.KeyMap

	Submit,
	Focus,
	Blur key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Submit,
		k.Focus,
		k.Blur,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
