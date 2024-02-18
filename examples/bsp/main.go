package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/bsp"
	"github.com/tuiphy/cans/textinput"
	"github.com/tuiphy/cans/timer"
	"github.com/tuiphy/cans/viewport"
	"github.com/tuiphy/soda"
)

func run() error {
	state := bsp.New(bsp.WithBranch(bsp.Branch{
		Direction: bsp.DirectionHorizontal,
		Ratio:     bsp.RatioGolden,
		Left:      NewBorderedState("left", textinput.New()),
		Right: &bsp.Branch{
			Direction: bsp.DirectionVertical,
			Ratio:     bsp.RatioGolden,
			Left: NewBorderedState(
				"right-left",
				timer.New(time.Second*100, timer.WithInterval(time.Millisecond)),
			),
			Right: NewBorderedState(
				"right-right",
				viewport.New("Hello, World!"),
			),
		},
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
