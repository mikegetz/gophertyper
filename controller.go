package main

import (
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
		for i, gopher := range m.gophers {
			if gopher.Y > 0 {
				m.gophers[i].Y--
			}
		}
		return m, moveGophers(3 * time.Second)

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}
	}

	if m.height > 0 && len(m.gophers) == 0 {
		m.gophers = []gopher{
			{X: 5, Y: m.height - 2},
			{X: 8, Y: m.height - 2},
			{X: 11, Y: m.height - 2},
		}
	}

	return m, nil
}
