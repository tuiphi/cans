package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbletea"
)

func New[I Item](items []I, options ...Option[I]) *State[I] {
	state := &State[I]{
		onSubmit: func(I) tea.Cmd {
			return nil
		},
		onMultiSubmit: func([]I) tea.Cmd {
			return nil
		},
		multiChoice:  false,
		itemHeight:   2,
		singularNoun: "item",
		pluralNoun:   "items",
		keyMap:       DefaultKeyMap(),
	}

	for _, option := range options {
		option(state)
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.FilterMatch.Underline(false)
	delegate.Styles.NormalTitle.Bold(true)
	delegate.Styles.SelectedTitle.Bold(true)
	delegate.SetHeight(state.itemHeight)

	if state.itemHeight == 1 {
		delegate.ShowDescription = false
	}

	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = &_Item[I]{
			internal: item,
		}
	}

	l := list.New(listItems, delegate, 0, 0)
	l.SetShowHelp(false)
	l.SetShowFilter(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowPagination(false)
	l.SetStatusBarItemName(state.singularNoun, state.pluralNoun)
	l.DisableQuitKeybindings()
	l.InfiniteScrolling = false
	l.KeyMap = state.keyMap.List
	l.Paginator.Type = paginator.Arabic

	state.list = l

	return state
}
