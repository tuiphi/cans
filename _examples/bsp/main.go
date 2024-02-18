package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/cans/bsp"
	"github.com/tuiphy/soda"
)

func run() error {
	// layout := bsp.New(bsp.NewTree(
	// 	bsp.NewLeaf(bordered.New(timer.New(time.Hour, timer.WithInterval(time.Millisecond)))),
	// 	bsp.NewTree(
	// 		bsp.NewLeaf(bordered.New(&SizeHolder{title: "right-left"})),
	// 		bsp.NewLeaf(bordered.New(&SizeHolder{title: "right-right"})),
	// 		bsp.WithTreeDirection(bsp.DirectionVertical),
	// 	),
	// ))
	state := bsp.New(bsp.WithBranch(bsp.Branch{
		Direction: bsp.DirectionHorizontal,
		Ratio:     bsp.RatioGolden,
		Left:      &SizeHolder{title: "left"},
		Right: &bsp.Branch{
			Direction: bsp.DirectionVertical,
			Ratio:     bsp.RatioGolden,
			Left:      &SizeHolder{title: "right-left"},
			Right:     &SizeHolder{title: "right-right"},
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
