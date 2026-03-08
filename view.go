package main

import (
	"fmt"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	gameContainerStyle        = lipgloss.NewStyle().Height(minTerminalHeight)
	containerStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#714209"))
	transitionContainerStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#422400"))
	waveTransitionTextStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#422400")).Foreground(lipgloss.Color("#d2d201")).Bold(true)
	gopherHoleSelectedStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#422400")).Foreground(lipgloss.Color("#d2d201")).Bold(true)
	skyStyle                  = lipgloss.NewStyle().Background(lipgloss.Color("#87a8eb")).Foreground(lipgloss.Black)
	grassySkyStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#009600")).Foreground(lipgloss.Color("#014a01"))
	grassyGroundStyle         = lipgloss.NewStyle().Background(lipgloss.Color("#4e2a00")).Foreground(lipgloss.Color("#228B22"))
	gopherHoleUnselectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#4e2a00")).Foreground(lipgloss.Color("#ffffff"))
)

func (m model) View() tea.View {
	screen := ""
	screen += m.printSky()

	if m.resizeWidthWarning || m.resizeHeightWarning {
		screen += m.printResizeWarning()
	} else if m.waveTransition {
		screen += m.printWaveTransition()
	} else if m.loseTransition {
		screen += m.printGophers(8)
		screen += m.printLoseTransition(8)
	} else if m.winTransition {
		screen += m.printWinTransition()
	} else {
		screen += m.printGophers(0)
	}

	// add height to m.topPadding to account for the ground and the newlines above
	//screen += m.debugView()

	teaView := tea.NewView(screen)
	return teaView
}

func (m model) printSky() string {
	screen := ""
	grassyGround := "~~^~^~^~~^~~^~*~^~^~~^~^~~^~"
	grassyGroundRepeats := (m.width / len(grassyGround)) + 1

	screen += skyStyle.Width(m.width).Render("[Esc] quit    [Space] pause/resume") + "\n"
	//screen += skyStyle.Width(m.width).Render("version: "+Version) + "\n"
	sky := skyStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	var horizon string
	if m.lose != nil {
		horizon = grassySkyStyle.Width(m.width).Render(strings.Repeat(" ", m.lose.X)+gopherHoleUnselectedStyle.Render("🐹")) + "\n"
	} else {
		padding := m.width - lipgloss.Width(Version)
		if padding < 0 {
			padding = 0
		}

		horizon = grassySkyStyle.Width(m.width).Render(strings.Repeat(" ", padding)+Version) + "\n"
	}
	// 3 accounts for top line controls, horizon, and ground
	for i := 1; i < m.topPadding-3; i++ {
		if m.lose != nil && i == m.topPadding-4 {
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

func (m model) printReport() string {
	correctKeypresses := float64(m.correctKeypresses) / float64(m.keypresses)
	m.accuracy = fmt.Sprintf("%.2f%%", correctKeypresses*100)
	report := strings.Repeat("\n", 5)
	report += "Gophers Per Minute (GPM): " + m.gpm + "\n"
	report += "Accuracy: " + m.accuracy + "\n"
	report += "Correct Keypresses: " + fmt.Sprint(m.correctKeypresses) + "/" + fmt.Sprint(m.keypresses) + "\n"

	return report
}

func (m model) printLoseTransition(loseVerticalPadding int) string {
	screen := ""
	gameViewSize := minTerminalHeight - m.topPadding

	screen += waveTransitionTextStyle.Width(m.width).Align(lipgloss.Center).Render(m.printReport()) + "\n"

	for y := 0; y < gameViewSize-(loseVerticalPadding+lipgloss.Height(m.printReport())); y++ {
		screen += transitionContainerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	return screen
}

func (m model) printWinTransition() string {
	screen := ""
	winVerticalPadding := 4
	gameViewSize := minTerminalHeight - m.topPadding

	for y := 0; y < winVerticalPadding; y++ {
		screen += transitionContainerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	screen += waveTransitionTextStyle.Width(m.width).Align(lipgloss.Center).Render(winText) + "\n"

	screen += waveTransitionTextStyle.Width(m.width).Align(lipgloss.Center).Render(m.printReport()) + "\n"

	for y := 0; y < gameViewSize-(winVerticalPadding+lipgloss.Height(m.printReport())+lipgloss.Height(wave)); y++ {
		screen += transitionContainerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	return screen
}

func (m model) printResizeWarning() string {
	screen := ""
	verticalPadding := 4
	gameViewSize := minTerminalHeight - m.topPadding

	for y := 0; y < verticalPadding; y++ {
		screen += containerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	if m.resizeWidthWarning {
		warning := "RESIZE TERMINAL to Minimum Width: " + fmt.Sprint(minTerminalWidth)
		screen += containerStyle.Width(m.width).Align(lipgloss.Center).Render(warning) + "\n"
	}

	if m.resizeHeightWarning {
		warning := "RESIZE TERMINAL to Minimum Height: " + fmt.Sprint(minTerminalHeight)
		screen += containerStyle.Width(m.width).Align(lipgloss.Center).Render(warning) + "\n"
	}

	for y := 0; y < (gameViewSize - (verticalPadding + 1)); y++ {
		screen += containerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	return screen
}

func (m model) printWaveTransition() string {
	screen := ""
	verticalPadding := 4
	gameViewSize := minTerminalHeight - m.topPadding

	for y := 0; y < verticalPadding; y++ {
		screen += transitionContainerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	screen += waveTransitionTextStyle.Width(m.width).Align(lipgloss.Center).Render(concatArt(wave, waveNumbers[m.wave])) + "\n"

	for y := 0; y < gameViewSize-(verticalPadding+lipgloss.Height(wave)); y++ {
		screen += transitionContainerStyle.Width(m.width).Render(strings.Repeat(" ", m.width)) + "\n"
	}

	return screen
}

func concatArt(left, right string) string {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")
	var result []string
	for i := 0; i < max(len(leftLines), len(rightLines)); i++ {
		l, r := "", ""
		if i < len(leftLines) {
			l = leftLines[i]
		}
		if i < len(rightLines) {
			r = rightLines[i]
		}
		result = append(result, l+r)
	}
	return strings.Join(result, "\n")
}

func (m model) printGophers(truncate int) string {
	gameViewSize := minTerminalHeight - m.topPadding

	if truncate > 0 {
		gameViewSize = truncate
	}

	screen := ""

	for y := 0; y < gameViewSize; y++ {
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
					} else if sortedGopher.Alive == false {
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
	debug := "[DEBUG]"

	//debug += fmt.Sprintf("Terminal Size: %d x %d", m.width, m.height)
	//debug += " Gophers: " + fmt.Sprint(m.gophers)
	//debug += fmt.Sprintf(" lose: %v", m.lose)
	//debug += fmt.Sprintf(" timeMultiplier: %d", m.timeMultiplier)
	return debug
}
