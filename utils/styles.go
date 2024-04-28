package utils

import "github.com/charmbracelet/lipgloss"

const (
	Blue   = lipgloss.Color("#B97AEE")
	Green  = lipgloss.Color("#9DF7E5")
	Orange = lipgloss.Color("#F4E4BA")
)

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Underline(true)

func ProjectColor(text string) string {
	return lipgloss.NewStyle().Foreground(Blue).Render(text)
}

func TimeColor(text string) string {
	return lipgloss.NewStyle().Foreground(Green).Render(text)
}

func TagColor(text string) string {
	return lipgloss.NewStyle().Foreground(Orange).Render(text)
}

func Faint(text string) string {
	return lipgloss.NewStyle().Faint(true).Render(text)
}
