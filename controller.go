package main

import (
	"math/rand"
	"time"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		randomGopher := rand.Intn(len(m.gophers))
		timeMultiplier := 10 - m.wave

		if m.gophers[randomGopher].Y > 0 {
			m.gophers[randomGopher].Y--
		} else {
			m.lose = &m.gophers[randomGopher]
			return m, nil
		}

		return m, moveGophers(time.Millisecond * time.Duration(timeMultiplier))

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}
	}

	initGophers(&m)

	return m, nil
}

func initGophers(m *model) {
	gopherCount := 10

	if m.height > 0 && len(m.gophers) == 0 {
		segmentWidth := m.width / gopherCount
		for i := 0; i < gopherCount; i++ {
			segmentStart := i * segmentWidth                                       // left edge of this gopher's segment
			segmentMargin := 1                                                     // columns reserved on each side to prevent adjacency
			usableWidth := segmentWidth - (segmentMargin * 2)                      // placeable range within the segment after margins
			gopherSpacing := segmentStart + segmentMargin + rand.Intn(usableWidth) // final X: start + margin + random offset
			m.gophers = append(m.gophers, gopher{X: gopherSpacing, Y: (m.height - m.topPadding), Word: easyWordList[rand.Intn(len(easyWordList))]})
		}
	}
}
