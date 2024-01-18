package bsp

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/cans/mux/window"
	"github.com/tuiphy/soda"
)

var _ Node = (*Tree)(nil)

type TreeOption func(*Tree)

func WithTreeDirection(direction Direction) TreeOption {
	return func(tree *Tree) {
		tree.direction = direction
	}
}

func WithTreeSplitRatio(ratio Ratio) TreeOption {
	return func(tree *Tree) {
		tree.splitRatio = ratio
	}
}

func NewTree(left, right Node, options ...TreeOption) *Tree {
	tree := &Tree{
		left:       left,
		right:      right,
		splitRatio: RatioEqual,
		direction:  DirectionHorizontal,
	}

	for _, option := range options {
		option(tree)
	}

	return tree
}

type Tree struct {
	left, right Node
	splitRatio  Ratio
	direction   Direction
}

// View implements Node.
func (t *Tree) View(layout soda.Layout) string {
	var join func(views ...string) string

	switch t.direction {
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

	return join(t.left.View(layout), t.right.View(layout))
}

// SetSize implements Node.
func (t *Tree) SetSize(size soda.Size) tea.Cmd {
	left, right := t.direction.Split(t.splitRatio, size)

	return tea.Batch(
		t.left.SetSize(left),
		t.right.SetSize(right),
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

	handleNode(t.left)
	handleNode(t.right)

	return states
}
