package utils

import "github.com/charmbracelet/lipgloss"

const (
	Purple = lipgloss.Color("#B97AEE")
	Green  = lipgloss.Color("#00A6A3")
	Yellow = lipgloss.Color("#F2B263")
)

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(Purple)

func ProjectColor(text string) string {
	return lipgloss.NewStyle().Foreground(Purple).Render(text)
}

func TimeColor(text string) string {
	return lipgloss.NewStyle().Foreground(Green).Render(text)
}

func TagColor(text string) string {
	return lipgloss.NewStyle().Foreground(Yellow).Render(text)
}
