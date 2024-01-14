package mux

import "github.com/charmbracelet/lipgloss"

func DefaultStyles() Styles {
	return Styles{
		ActiveBorderColor:  "#DCDCDC",
		DefaultBorderColor: "#A9A9A9",
	}
}

type Styles struct {
	ActiveBorderColor,
	DefaultBorderColor lipgloss.Color
}
