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
		if m.gophers[randomGopher].Y > 0 {
			m.gophers[randomGopher].Y--
		}
		return m, moveGophers(time.Millisecond * 250)

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
			gopherSpacing := (i * segmentWidth) + rand.Intn(segmentWidth)
			m.gophers = append(m.gophers, gopher{X: gopherSpacing, Y: (m.height - m.topPadding), Word: easyWordList[rand.Intn(len(easyWordList))]})
		}
	}
}
