package main

import (
	"fmt"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	containerStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#714209"))
	gopherHoleSelectedStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#422400")).Foreground(lipgloss.Color("#13ce32")).Bold(true)
	skyStyle                  = lipgloss.NewStyle().Background(lipgloss.Color("#87a8eb")).Foreground(lipgloss.Black)
	grassySkyStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#009600")).Foreground(lipgloss.Color("#014a01"))
	grassyGroundStyle         = lipgloss.NewStyle().Background(lipgloss.Color("#714209")).Foreground(lipgloss.Color("#228B22"))
	gopherHoleUnselectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#4e2a00"))
)

func (m model) View() tea.View {
	screen := ""
	screen += m.printSky()

	screen += m.printGophers()
	// add height to m.topPadding to account for the ground and the newlines above
	//screen += m.debugView()

	teaView := tea.NewView(screen)
	return teaView
}

func (m model) printSky() string {
	screen := ""
	grassyGround := "~~^~^~^~~^~~^~*~^~^~~^~^~~^~"
	grassyGroundRepeats := (m.width / len(grassyGround)) + 1

	screen += skyStyle.Width(m.width).Render("ctrl+c/esc to quit") + "\n"
	sky := skyStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	var horizon string
	if m.lose != nil {
		horizon = grassySkyStyle.Width(m.width).Render(strings.Repeat(" ", m.lose.X)+gopherHoleUnselectedStyle.Render("🐹")) + "\n"
	} else {
		horizon = grassySkyStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}
	for i := 1; i < m.topPadding-4; i++ {
		if m.lose != nil && i == m.topPadding-5 {
			loseText := "you lose to gopher "
			padding := m.lose.X - len(loseText)
			if padding < 0 {
				padding = 0
			}
			screen += skyStyle.Width(m.width).Render(strings.Repeat(" ", padding)+loseText+m.lose.Word) + "\n"
		} else {
			screen += sky
		}
	}

	screen += horizon

	ground := grassyGroundStyle.Width(m.width).Render(strings.Repeat(grassyGround, grassyGroundRepeats)[:m.width]) + "\n"

	screen += ground

	return screen
}

func (m model) printGophers() string {
	screen := ""

	for y := 0; y < (m.height - m.topPadding); y++ {
		sortedLineGophers := []gopher{}
		for _, gopher := range m.gophers {
			if gopher.Y == y {
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
				gopherHoleStyle := gopherHoleUnselectedStyle
				if m.selected != nil && sortedGopher.X == m.selected.X {
					gopherHoleStyle = gopherHoleSelectedStyle
				}
				var renderObject string
				switch sortedGopher.Type {
				case gopherIcon:
					if m.lose != nil && sortedGopher.X == m.lose.X {
						renderObject = gopherHoleStyle.Render("  ")
					}
					if sortedGopher.Alive == false {
						renderObject = "💀"
					} else {
						renderObject = "🐹"
					}
				case gopherPath:
					if (y - sortedGopher.Y) <= len(sortedGopher.DisplayWord) {
						renderObject = gopherHoleStyle.Render(string(sortedGopher.DisplayWordRunes()[(y-sortedGopher.Y)-1]) + " ")
					} else {
						renderObject = gopherHoleStyle.Render("  ")
					}
				}

				padding := sortedGopher.X - lineOffset

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

	return screen
}

func (m model) debugView() string {
	debug := "[DEBUG] "

	debug += fmt.Sprintf("Terminal Size: %d x %d", m.width, m.height)
	debug += " Gophers: " + fmt.Sprint(m.gophers)
	return debug
}
