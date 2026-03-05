package main

import (
	"fmt"
	"slices"
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
		sortedLineGophers := []gopher{}
		for _, gopher := range m.gophers {
			if gopher.Y == y+1 {
				gopher.Type = word
				sortedLineGophers = append(sortedLineGophers, gopher)
			} else if gopher.Y == y {
				gopher.Type = gopherIcon
				sortedLineGophers = append(sortedLineGophers, gopher)
			} else if gopher.Y < y {
				gopher.Type = gopherPath
				sortedLineGophers = append(sortedLineGophers, gopher)
			}
		}
		slices.SortFunc(sortedLineGophers, func(i, j gopher) int {
			return i.X - j.X
		})

		if len(sortedLineGophers) > 0 {
			var line string
			lineOffset := 0
			for _, sortedGopher := range sortedLineGophers {

				var renderObject string
				switch sortedGopher.Type {
				case word:
					renderObject = sortedGopher.Word
				case gopherIcon:
					renderObject = "🐹"
				case gopherPath:
					renderObject = gopherHoleStyle.Render("  ")
				}

				padding := sortedGopher.X - lineOffset

				if sortedGopher.Type == word {
					padding -= lipgloss.Width(renderObject) / 2
				}

				if padding < 0 {
					padding = 0
				}

				line += containerStyle.Render(strings.Repeat(" ", padding) + renderObject)
				lineOffset += padding + lipgloss.Width(renderObject)
			}
			screen += containerStyle.Width(m.width).Render(line) + "\n"
		} else {
			screen += containerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
		}
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
