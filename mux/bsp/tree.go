package bsp

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/cans/mux/window"
	"github.com/tuiphy/soda"
)

var _ Node = (*Tree)(nil)

type Tree struct {
	Left, Right Node
	SplitRatio  Ratio
	Direction   Direction
}

// View implements Node.
func (t *Tree) View(layout soda.Layout) string {
	var join func(views ...string) string

	switch t.Direction {
	case DirectionVertical:
		join = func(views ...string) string {
			return lipgloss.JoinVertical(lipgloss.Left, views...)
		}
	case DirectionHorizontal:
		join = func(views ...string) string {
			return lipgloss.JoinHorizontal(lipgloss.Left, views...)
		}
	default:
		panic("unreachable")
	}

	return join(t.Left.View(layout), t.Right.View(layout))
}

// SetSize implements Node.
func (t *Tree) SetSize(size soda.Size) tea.Cmd {
	left, right := t.Direction.Split(t.SplitRatio, size)

	return tea.Batch(
		t.Left.SetSize(left),
		t.Right.SetSize(right),
	)
}

func (t *Tree) Windows() []window.Window {
	var states []window.Window

	handleNode := func(node Node) {
		switch node := node.(type) {
		case *Tree:
			states = append(states, node.Windows()...)
		case *Leaf:
			states = append(states, node)
		}
	}

	handleNode(t.Left)
	handleNode(t.Right)

	return states
}
