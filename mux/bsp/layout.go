package bsp

import "github.com/tuiphy/cans/mux/window"

func New(root Node) *Layout {
	return &Layout{
		Node: root,
	}
}

type Layout struct {
	Node
}

func (l *Layout) Windows() []window.Window {
	switch node := l.Node.(type) {
	case *Tree:
		return node.Windows()
	case *Leaf:
		return []window.Window{node}
	default:
		panic("unreachable")
	}
}
