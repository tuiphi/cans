package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/tuiphy/cans/filepicker"
	"github.com/tuiphy/cans/viewport"
	"github.com/tuiphy/soda"
	"log"
	"os"
)

func run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filePickerKeyMap := filepicker.DefaultKeyMap()
	filePickerKeyMap.Back = key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "back"),
	)

	model := soda.New(filepicker.New(
		filepicker.WithKeyMap(filePickerKeyMap),
		filepicker.WithDir(home),
		filepicker.WithOnSelect(func(path string) tea.Cmd {
			return soda.Wrap(func() tea.Cmd {
				contents, err := os.ReadFile(path)
				if err != nil {
					return soda.SendError(err)
				}

				return soda.PushState(
					viewport.New(
						string(contents),
						viewport.WithResizeContent(func(content string, size soda.Size) string {
							return wrap.String(content, size.Width)
						}),
					),
				)
			})
		}),
	))

	program := tea.NewProgram(model, tea.WithAltScreen())
	_, err = program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
