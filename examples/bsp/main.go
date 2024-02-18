package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"github.com/tuiphy/cans/bsp"
	"github.com/tuiphy/cans/textinput"
	"github.com/tuiphy/cans/timer"
	"github.com/tuiphy/cans/viewport"
	"github.com/tuiphy/soda"
)

const dummyText = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel libero nulla. Donec eget libero elementum, mollis erat eu, malesuada diam. Donec maximus molestie pharetra. Suspendisse dignissim nisi vitae nunc ultricies cursus. Phasellus semper nisl augue, nec faucibus orci accumsan id. Quisque vitae lectus tempus, porta turpis vitae, cursus eros. Integer non augue condimentum purus tincidunt ultricies eu vel mauris. Morbi non efficitur lectus. Maecenas magna augue, pharetra id consequat ac, vestibulum id nisi. Curabitur congue nibh a neque lobortis, eget dignissim magna consectetur.

Duis gravida risus eu placerat ultrices. Cras pulvinar, nisl ut malesuada vestibulum, elit felis dapibus lorem, ut iaculis orci ex eget dolor. Mauris non nunc libero. Maecenas a sodales neque. Vestibulum at ligula ac nisl luctus luctus sed ut metus. Etiam vitae neque id est dapibus vulputate. In id orci dolor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus id semper enim. Suspendisse potenti. Praesent molestie elit sit amet augue maximus egestas. Sed imperdiet eget ligula ac consequat. Donec eleifend lacus ex, ut accumsan felis vehicula sit amet. Phasellus at eros ut dui dignissim ornare. Nullam iaculis nibh eget mollis porttitor.

Phasellus feugiat orci massa, vitae ultrices felis imperdiet sed. Morbi eget ex luctus, semper risus vel, tristique elit. Sed quam enim, interdum congue vehicula non, scelerisque id est. Nam commodo condimentum lobortis. Nunc at dui dolor. Proin eu varius mauris. Nunc volutpat facilisis urna non dapibus.

Mauris posuere justo quis mi imperdiet suscipit sit amet ut nisl. Nam interdum, purus at mollis pretium, arcu mauris rhoncus leo, eget viverra dolor dui sit amet diam. Curabitur laoreet non ante non tincidunt. Mauris efficitur posuere magna, id cursus velit tincidunt eu. Donec feugiat nibh quis tellus laoreet dignissim. Nullam viverra tincidunt urna. Nulla tincidunt erat quis tortor facilisis luctus. Praesent commodo laoreet metus facilisis finibus. Maecenas est lorem, posuere id vulputate ut, tempor ac lorem. Duis blandit quis tellus viverra sagittis. Donec hendrerit mauris at scelerisque vestibulum. Sed finibus efficitur porta. Praesent vel erat feugiat, rutrum tortor fringilla, pretium nibh. Aenean dignissim porta velit id fermentum. Integer congue vitae mi ac lobortis.

Nam sit amet justo arcu. Ut blandit ex est, eget accumsan dolor posuere venenatis. Sed vel facilisis enim, vel venenatis metus. Suspendisse mollis eget libero eu ornare. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Sed in lectus et neque scelerisque efficitur. Fusce sodales arcu id suscipit tincidunt. Pellentesque laoreet erat sed orci consectetur mattis. Etiam sapien quam, posuere viverra gravida congue, dictum at elit. Donec at cursus lorem. Cras non fringilla lectus, a viverra ligula. Quisque posuere sagittis nibh, nec suscipit turpis accumsan eget.`

func run() error {
	state := bsp.New(bsp.WithCycle(true), bsp.WithBranch(bsp.Branch{
		Direction: bsp.DirectionHorizontal,
		Ratio:     bsp.RatioEven,
		Left:      NewBorderedState("Input", textinput.New()),
		Right: &bsp.Branch{
			Direction: bsp.DirectionVertical,
			Ratio:     bsp.RatioEven,
			Left: NewBorderedState(
				"Timer",
				timer.New(time.Second*100, timer.WithInterval(time.Millisecond)),
			),
			Right: NewBorderedState(
				"Viewport",
				viewport.New(dummyText, viewport.WithResizeContent(func(content string, size soda.Size) string {
					return wordwrap.String(content, size.Width)
				})),
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
