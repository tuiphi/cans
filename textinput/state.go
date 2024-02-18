package textinput

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	textInput textinput.Model

	suggestionsSupplier SuggestionsSupplier

	keyMap   KeyMap
	onSubmit OnSubmitFunc
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return s.textInput.Focused()
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	promptWidth := lipgloss.Width(s.textInput.PromptStyle.Render(s.textInput.Prompt))
	s.textInput.Width = size.Width - promptWidth - 1
	return nil
}

func (s *State) Title() string {
	return "Text input"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	return tea.Batch(s.focus(), textinput.Blink)
}

func (s *State) focus() tea.Cmd {
	s.keyMap.Focus.SetEnabled(false)
	s.keyMap.Blur.SetEnabled(true)
	s.keyMap.Submit.SetEnabled(true)
	return s.textInput.Focus()
}

func (s *State) blur() tea.Cmd {
	s.keyMap.Focus.SetEnabled(true)
	s.keyMap.Blur.SetEnabled(false)
	s.keyMap.Submit.SetEnabled(false)
	s.textInput.Blur()
	return nil
}

func (s *State) toggleFocus() tea.Cmd {
	if s.Focused() {
		return s.blur()
	}

	return s.focus()
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Submit):
			return s.onSubmit(s.textInput.Value())
		case key.Matches(msg, s.keyMap.Focus, s.keyMap.Blur):
			return s.toggleFocus()
		}
	}

	var cmd tea.Cmd
	s.textInput, cmd = s.textInput.Update(msg)
	s.textInput.SetSuggestions(s.suggestionsSupplier(s.textInput.Value()))
	return cmd
}

func (s *State) View(soda.Layout) string {
	return s.textInput.View()
}
