package list

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
	"strconv"
	"strings"
)

var _ soda.State = (*State[list.DefaultItem])(nil)

type State[I Item] struct {
	list list.Model

	onSubmit      OnSubmitFunc[I]
	onMultiSubmit OnMultiSubmitFunc[I]

	itemHeight int

	singularNoun, pluralNoun string

	multiChoice bool

	keyMap KeyMap
}

func (s *State[I]) Destroy() {
}

func (s *State[I]) Focused() bool {
	return s.list.FilterState() != list.Unfiltered
}

func (s *State[I]) SetSize(size soda.Size) tea.Cmd {
	s.list.SetSize(size.Width, size.Height)

	if s.list.Paginator.TotalPages > 5 {
		s.list.Paginator.Type = paginator.Arabic
	} else {
		s.list.Paginator.Type = paginator.Dots
	}

	return nil
}

func (s *State[I]) Title() string {
	return "List"
}

func (s *State[I]) Subtitle() string {
	singular, plural := s.list.StatusBarItemName()

	var subtitle strings.Builder

	subtitle.Grow(max(len(singular), len(plural)) * 2)

	itemsCount := len(s.list.VisibleItems())
	subtitle.WriteString(strconv.Itoa(itemsCount))
	subtitle.WriteString(" ")

	if itemsCount == 1 {
		subtitle.WriteString(singular)
	} else {
		subtitle.WriteString(plural)
	}

	if s.list.FilterState() == list.FilterApplied {
		subtitle.WriteString(" ")
		subtitle.WriteString(fmt.Sprintf("%q", s.list.FilterValue()))
	}

	return subtitle.String()
}

func (s *State[I]) Layout() (layout soda.Layout, override bool) {
	return
}

func (s *State[I]) Status() string {
	if s.list.FilterState() == list.Filtering {
		return s.list.FilterInput.View()
	}

	if s.list.Paginator.TotalPages > 1 {
		return s.list.Paginator.View()
	}

	return ""
}

func (s *State[I]) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State[I]) Init(ctx context.Context) tea.Cmd {
	if !s.multiChoice {
		s.keyMap.Select.SetEnabled(false)
	}

	return nil
}

func (s *State[I]) Items() []I {
	listItems := s.list.Items()

	items := make([]I, 0, cap(listItems))
	for _, listItem := range listItems {
		item, ok := listItem.(*_Item[I])
		if ok {
			items = append(items, item.internal)
		}
	}

	return items
}

func (s *State[I]) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s.list.FilterState() == list.Filtering {
			goto updateList
		}

		switch {
		case key.Matches(msg, s.keyMap.Select):
			item, ok := s.list.SelectedItem().(*_Item[I])
			if !ok {
				return nil
			}

			item.ToggleSelected()
			return nil
		case key.Matches(msg, s.keyMap.Submit):
			if s.multiChoice {
				return s.submitMulti()
			}

			return s.submitSingle()
		}
	}

updateList:
	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	return cmd
}

func (s *State[I]) submitMulti() tea.Cmd {
	listItems := s.list.Items()

	submitItems := make([]I, 0, cap(listItems))
	for _, listItem := range listItems {
		item, ok := listItem.(*_Item[I])
		if ok && item.selected {
			submitItems = append(submitItems, item.internal)
		}
	}

	return s.onMultiSubmit(submitItems)
}

func (s *State[I]) submitSingle() tea.Cmd {
	item, ok := s.list.SelectedItem().(*_Item[I])
	if !ok {
		return nil
	}

	return s.onSubmit(item.internal)
}

func (s *State[I]) View() string {
	return s.list.View()
}
