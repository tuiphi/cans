package main

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/cans/bordered"
	"github.com/tuiphy/cans/bsp"
	"github.com/tuiphy/soda"
)

var (
	_ soda.State = (*Bordered)(nil)
	_ bsp.Leaf   = (*Bordered)(nil)
)

func NewBorderedState(title string, inner soda.State) *Bordered {
	return &Bordered{
		title:         title,
		activeColor:   lipgloss.Color("1"),
		inactiveColor: lipgloss.Color("4"),
		bordered:      bordered.New(inner),
	}
}

type Bordered struct {
	title string

	activeColor, inactiveColor lipgloss.Color

	bordered *bordered.State
}

// Activate implements bsp.Activater.
func (d *Bordered) Activate() tea.Cmd {
	d.bordered.SetForeground(d.activeColor)
	return nil
}

// Deactivate implements bsp.Activater.
func (d *Bordered) Deactivate() tea.Cmd {
	d.bordered.SetForeground(d.inactiveColor)
	return nil
}

// Destroy implements soda.State.
func (b *Bordered) Destroy() {
	b.bordered.Destroy()
}

// Focused implements soda.State.
func (b *Bordered) Focused() bool {
	return b.bordered.Focused()
}

// Init implements soda.State.
func (d *Bordered) Init(ctx context.Context) tea.Cmd {
	d.bordered.SetForeground(d.inactiveColor)
	return d.bordered.Init(ctx)
}

// KeyMap implements soda.State.
func (b *Bordered) KeyMap() help.KeyMap {
	return b.bordered.KeyMap()
}

// Layout implements soda.State.
func (*Bordered) Layout() (layout soda.Layout, override bool) {
	return
}

// SetSize implements soda.State.
func (d *Bordered) SetSize(size soda.Size) tea.Cmd {
	return d.bordered.SetSize(size)
}

// Status implements soda.State.
func (b *Bordered) Status() string {
	return b.bordered.Status()
}

// Subtitle implements soda.State.
func (b *Bordered) Subtitle() string {
	return b.bordered.Status()
}

// Title implements soda.State.
func (d *Bordered) Title() string {
	return d.title
}

// Update implements soda.State.
func (b *Bordered) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	return b.bordered.Update(ctx, msg)
}

// View implements soda.State.
func (d *Bordered) View(layout soda.Layout) string {
	return d.bordered.View(layout)
}
