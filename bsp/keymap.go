package bsp

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Next: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next"),
		),
		Prev: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("backtab", "prev"),
		),
	}
}

var _ help.KeyMap = (*KeyMap)(nil)

type KeyMap struct {
	Next, Prev key.Binding
}

// FullHelp implements help.KeyMap.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		append(k.ShortHelp(), k.Prev),
	}
}

// ShortHelp implements help.KeyMap.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Next,
	}
}

func (k KeyMap) with(other help.KeyMap) combinedKeyMap {
	return combinedKeyMap{
		overlays: []help.KeyMap{k, other},
	}
}

var _ help.KeyMap = (*combinedKeyMap)(nil)

type combinedKeyMap struct {
	overlays []help.KeyMap
}

func (c combinedKeyMap) ShortHelp() []key.Binding {
	var bindings []key.Binding

	for _, overlay := range c.overlays {
		bindings = append(bindings, overlay.ShortHelp()...)
	}

	return bindings
}

func (c combinedKeyMap) FullHelp() [][]key.Binding {
	var groups [][]key.Binding

	for _, overlay := range c.overlays {
		groups = append(groups, overlay.FullHelp()...)
	}

	return groups
}

func (c combinedKeyMap) With(other help.KeyMap) combinedKeyMap {
	return combinedKeyMap{
		overlays: append(c.overlays, other),
	}
}
