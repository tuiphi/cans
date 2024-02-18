package bsp

import tea "github.com/charmbracelet/bubbletea"

type Leaf interface {
	Activate() tea.Cmd
	Deactivate() tea.Cmd
}
