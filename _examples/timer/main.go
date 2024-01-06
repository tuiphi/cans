package main

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/textinput"
	"github.com/tuiphy/cans/timer"
	"github.com/tuiphy/soda"
	"log"
	"time"
)

func run() error {
	state := textinput.New(textinput.WithOnSubmit(func(value string) tea.Cmd {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return soda.SendError(err)
		}

		return soda.PushState(timer.New(
			duration,
			timer.WithInterval(time.Millisecond),
			timer.WithAutoStart(false),
			timer.WithOnTimeout(soda.Wrap(func() tea.Cmd {
				return soda.Notify("Timeout")
			})),
		))
	}))

	model := soda.New(state)
	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
