package bsp

import tea "github.com/charmbracelet/bubbletea"

type Activater interface {
	Activate() tea.Cmd
	Deactivate() tea.Cmd
}
