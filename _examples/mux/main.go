package main

import (
	"log"
	"time"

	"github.com/tuiphy/cans/bordered"

	"github.com/tuiphy/cans/timer"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/mux"
	"github.com/tuiphy/cans/mux/bsp"
	"github.com/tuiphy/soda"
)

func run() error {
	layout := bsp.New(bsp.NewTree(
		bsp.NewLeaf(bordered.New(timer.New(time.Hour, timer.WithInterval(time.Millisecond)))),
		bsp.NewTree(
			bsp.NewLeaf(bordered.New(&SizeHolder{title: "right-left"})),
			bsp.NewLeaf(bordered.New(&SizeHolder{title: "right-right"})),
			bsp.WithTreeDirection(bsp.DirectionVertical),
		),
	))

	state := mux.New(layout)
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
