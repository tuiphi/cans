package mux

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		ActivateNext: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next"),
		),
		ActivatePrev: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("backtab", "prev"),
		),
	}
}

type KeyMap struct {
	ActivateNext, ActivatePrev key.Binding
}

// FullHelp implements help.KeyMap.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.ActivateNext,
			k.ActivatePrev,
		},
	}
}

// ShortHelp implements help.KeyMap.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
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
		if overlay != nil {
			bindings = append(bindings, overlay.ShortHelp()...)
		}
	}

	return bindings
}

func (c combinedKeyMap) FullHelp() [][]key.Binding {
	var groups [][]key.Binding

	for _, overlay := range c.overlays {
		if overlay != nil {
			groups = append(groups, overlay.FullHelp()...)
		}
	}

	return groups
}

func (c combinedKeyMap) With(other help.KeyMap) combinedKeyMap {
	return combinedKeyMap{
		overlays: append(c.overlays, other),
	}
}
