package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	containerStyle = lipgloss.NewStyle()
)

func (m model) View() tea.View {
	screen := ""
	debug := "[DEBUG] "

	debug += fmt.Sprintf("Terminal Size: %d x %d", m.width, m.height)
	debug += " Gophers: " + fmt.Sprint(m.gophers)

	for y := 0; y < m.height; y++ {
		for _, gopher := range m.gophers {
			if gopher.Y == y {
				screen += containerStyle.Width(m.width).Render(strings.Repeat(" ", gopher.X) + "🐹\n")
			} else {
				screen += "\n"
			}
		}
	}

	screen += debug

	teaView := tea.NewView(screen)
	return teaView
}
