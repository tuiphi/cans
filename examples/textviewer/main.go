package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
	"log"
	"os"
)

func run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	state := New(home)

	modelKeyMap := soda.DefaultKeyMap()
	modelKeyMap.Back.SetKeys("backspace")
	modelKeyMap.Back.SetHelp("backspace", "back")

	model := soda.New(state, soda.WithKeyMap(modelKeyMap))

	program := tea.NewProgram(model, tea.WithAltScreen())
	_, err = program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
