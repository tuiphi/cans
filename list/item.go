package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Item = list.DefaultItem

type _Item[I Item] struct {
	internal I
	selected bool
}

func (i *_Item[I]) Title() string {
	if i.selected {
		return i.internal.Title() + " " + lipgloss.NewStyle().Bold(true).Render("•")
	}

	return i.internal.Title()
}

func (i *_Item[I]) FilterValue() string {
	return i.internal.FilterValue()
}

func (i *_Item[I]) Description() string {
	return i.internal.Description()
}

func (i *_Item[I]) ToggleSelected() {
	i.selected = !i.selected
}
