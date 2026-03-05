package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	containerStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#714209"))
	skyStyle          = lipgloss.NewStyle().Background(lipgloss.Color("#87a8eb")).Foreground(lipgloss.Black)
	grassySkyStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#009600")).Foreground(lipgloss.Color("#014a01"))
	grassyGroundStyle = lipgloss.NewStyle().Background(lipgloss.Color("#714209")).Foreground(lipgloss.Color("#228B22"))
	gopherHoleStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#422400"))
)

func (m model) View() tea.View {
	screen := ""
	grassyGround := "~~^~^~^~~^~~^~*~^~^~~^~^~~^~"
	grassyGroundRepeats := (m.width / len(grassyGround)) + 1

	screen += skyStyle.Width(m.width).Render("ctrl+c/esc to quit") + "\n"
	sky := skyStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	skyLast := grassySkyStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	for i := 1; i < m.topPadding-4; i++ {
		screen += sky
	}

	screen += skyLast

	ground := grassyGroundStyle.Width(m.width).Render(strings.Repeat(grassyGround, grassyGroundRepeats)[:m.width]) + "\n"

	screen += ground
	for y := 0; y < (m.height - m.topPadding); y++ {
		for _, gopher := range m.gophers {
			if gopher.Y == y+1 {
				wordX := gopher.X - (len(gopher.Word) / 2)
				if wordX < 0 {
					wordX = 0
				}
				screen += containerStyle.Render(strings.Repeat(" ", wordX) + gopher.Word)
			}
			if gopher.Y == y {
				screen += containerStyle.Render(strings.Repeat(" ", gopher.X) + "🐹")
			}
			if gopher.Y < y {
				screen += containerStyle.Render(strings.Repeat(" ", gopher.X) + gopherHoleStyle.Render("  "))
			}
		}
		screen += containerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
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
