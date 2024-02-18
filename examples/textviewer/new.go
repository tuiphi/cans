package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/tuiphy/cans/filepicker"
	"github.com/tuiphy/cans/viewport"
	"github.com/tuiphy/soda"
	"os"
)

func New(dir string) *State {
	fp := filepicker.New(
		filepicker.WithDir(dir),
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
	)

	return &State{
		filePicker: fp,
	}
}
