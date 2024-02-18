package bsp

import tea "github.com/charmbracelet/bubbletea"

func SetBranch(branch Branch) tea.Cmd {
	return func() tea.Msg {
		return BranchSetMsg{
			Branch: branch,
		}
	}
}
