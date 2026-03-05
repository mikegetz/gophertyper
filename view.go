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
	screen := strings.Repeat("\n", m.topPadding-1)

	grassyGround := "~~^~^~^~~^~~^~*~^~^~~^~^~~^~"

	grassyGroundRepeats := (m.width / len(grassyGround)) + 1

	ground := containerStyle.Width(m.width).Render(strings.Repeat(grassyGround, grassyGroundRepeats)[:m.width]) + "\n"

	screen += ground
	for y := 0; y < (m.height - m.topPadding); y++ {
		for _, gopher := range m.gophers {
			if gopher.Y == y {
				screen += containerStyle.Render(strings.Repeat(" ", gopher.X) + "🐹")
			}
			if gopher.Y < y {
				screen += containerStyle.Render(strings.Repeat(" ", gopher.X) + "||")
			}
		}
		screen += "\n"
	}

	// add height to m.topPadding to account for the ground and the newlines above
	//screen += m.debugView()

	teaView := tea.NewView(screen)
	return teaView
}

func (m model) debugView() string {
	debug := "[DEBUG] "

	debug += fmt.Sprintf("Terminal Size: %d x %d", m.width, m.height)
	debug += " Gophers: " + fmt.Sprint(m.gophers)
	return debug
}
