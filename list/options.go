package list

import tea "github.com/charmbracelet/bubbletea"

type (
	OnSubmitFunc[I Item]      func(item I) tea.Cmd
	OnMultiSubmitFunc[I Item] func(items []I) tea.Cmd
)

type Option[I Item] func(*State[I])

func WithMultiChoice[I Item](multiChoice bool) Option[I] {
	return func(state *State[I]) {
		state.multiChoice = multiChoice
	}
}

func WithOnSubmit[I Item](submitFunc OnSubmitFunc[I]) Option[I] {
	return func(state *State[I]) {
		state.onSubmit = submitFunc
	}
}

func WithOnMultiSubmit[I Item](multiSubmitFunc OnMultiSubmitFunc[I]) Option[I] {
	return func(state *State[I]) {
		state.onMultiSubmit = multiSubmitFunc
	}
}

func WithItemHeight[I Item](height int) Option[I] {
	return func(state *State[I]) {
		state.itemHeight = height
	}
}

func WithItemName[I Item](singular, plural string) Option[I] {
	return func(state *State[I]) {
		state.singularNoun = singular
		state.pluralNoun = plural
	}
}
