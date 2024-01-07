package main

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/list"
	"github.com/tuiphy/cans/viewport"
	"github.com/tuiphy/soda"
	"log"
	"strings"
)

func run() error {
	state := list.New[Item](
		[]Item{
			{
				title:       "Pizza",
				description: "Pepperoni, Cheese, Sauce",
			},
			{
				title:       "Burger",
				description: "Beef, Cheese, Pickles, Sauce",
			},
		},
		list.WithMultiChoice[Item](true),
		list.WithOnMultiSubmit[Item](func(items []Item) tea.Cmd {
			var sb strings.Builder

			sb.Grow(30 * len(items))

			sb.WriteString("You selected:\n")

			for _, item := range items {
				fmt.Fprintf(&sb, "- %s\n", item.title)
			}

			return soda.PushState(viewport.New(sb.String()))
		}),
	)

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
