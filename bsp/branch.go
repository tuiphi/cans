package bsp

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*Branch)(nil)

type Branch struct {
	Direction   Direction
	Ratio       Ratio
	Left, Right soda.State

	size soda.Size
}

// Destroy implements soda.State.
func (b *Branch) Destroy() {
	b.Left.Destroy()
	b.Right.Destroy()
}

// Focused implements soda.State.
func (*Branch) Focused() bool {
	return false
}

// Init implements soda.State.
func (b *Branch) Init(ctx context.Context) tea.Cmd {
	return tea.Batch(
		b.Left.Init(ctx),
		b.Right.Init(ctx),
	)
}

// KeyMap implements soda.State.
func (*Branch) KeyMap() help.KeyMap {
	panic("unimplemented")
}

// Layout implements soda.State.
func (*Branch) Layout() (layout soda.Layout, override bool) {
	return
}

func (b *Branch) Sizes() (left, right soda.Size) {
	return b.Direction.Split(b.size, b.Ratio)
}

// SetSize implements soda.State.
func (b *Branch) SetSize(size soda.Size) tea.Cmd {
	b.size = size

	switch {
	case b.Left == nil && b.Right == nil:
		return nil
	case b.Left == nil:
		return b.Right.SetSize(size)
	case b.Right == nil:
		return b.Left.SetSize(size)
	default:
		left, right := b.Sizes()

		return tea.Batch(
			b.Left.SetSize(left),
			b.Right.SetSize(right),
		)
	}
}

// Status implements soda.State.
func (*Branch) Status() string {
	return ""
}

// Subtitle implements soda.State.
func (*Branch) Subtitle() string {
	return ""
}

// Title implements soda.State.
func (*Branch) Title() string {
	return ""
}

// Update implements soda.State.
func (b *Branch) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch {
	case b.Left == nil && b.Right == nil:
		return nil
	case b.Left == nil:
		return b.Right.Update(ctx, msg)
	case b.Right == nil:
		return b.Left.Update(ctx, msg)
	default:
		return tea.Batch(
			b.Left.Update(ctx, msg),
			b.Right.Update(ctx, msg),
		)
	}
}

// View implements soda.State.
func (b *Branch) View(layout soda.Layout) string {
	view := func(state soda.State, layout soda.Layout, size soda.Size) string {
		if branch, ok := state.(*Branch); ok {
			return branch.View(layout)
		}

		// for regular states we need to fill empty space
		if custom, override := state.Layout(); override {
			layout = custom
		}

		return lipgloss.Place(
			size.Width,
			size.Height,
			layout.Horizontal,
			layout.Vertical,
			state.View(layout),
		)
	}

	switch {
	case b.Left == nil && b.Right == nil:
		return ""
	case b.Left == nil:
		return view(b.Right, layout, b.size)
	case b.Right == nil:
		return view(b.Left, layout, b.size)
	default:
		leftSize, rightSize := b.Sizes()

		return b.Direction.Join(
			view(b.Left, layout, leftSize),
			view(b.Right, layout, rightSize),
		)
	}
}

func (b Branch) States() []soda.State {
	var states []soda.State

	collect := func(s soda.State) {
		if branch, ok := s.(*Branch); ok {
			states = append(states, branch.States()...)
			return
		}

		states = append(states, s)
	}

	collect(b.Left)
	collect(b.Right)

	return states
}
