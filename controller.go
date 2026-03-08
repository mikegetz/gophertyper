package main

import (
	"math/rand/v2"
	"slices"
	"time"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	waveMultiplier := 50 - m.wave
	livingGopherMultiplier := (10 - m.LivingGopherCount()) * 15
	terminalHeightMultiplier := func() int {
		if m.height > 40 {
			return 0
		}
		return (m.height - m.topPadding) * 5
	}()
	m.timeMultiplier = waveMultiplier + livingGopherMultiplier + terminalHeightMultiplier
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if msg.Width < minTerminalWidth {
			m.resizeWarning = true
			return m, nil
		} else if m.resizeWarning {
			m.resizeWarning = false
			return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))
		} else {
			return m, nil
		}

	case winTransitionMsg:
		return m, tea.Quit

	case loseTransitionMsg:
		return m, tea.Quit

	case waveTransitionMsg:
		m.waveTransition = false
		m.wave++
		m.clearGophers()
		m.initGophers()
		return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))

	case tickMsg:
		if m.resizeWarning || m.win != nil || m.lose != nil || m.pause {
			return m, nil
		}

		randomGopher := m.RandomLivingGopher()

		if randomGopher == nil {
			if m.wave > 9 {
				m.win = m.selected
				return m, winTransition(&m, time.Second*5)
			}
			return m, waveTransition(&m, time.Second*3)
		}

		if randomGopher.Y > 0 {
			randomGopher.Y--
		} else {
			m.lose = randomGopher
			return m, loseTransition(&m, time.Second*15)
		}

		if m.selected != nil {
			if len(m.selected.DisplayWord) == 0 {
				m.selected.Alive = false
				m.selected = nil
			}
		}

		return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))

	case tea.KeyPressMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, m.keys.Pause) {
			m.pause = !m.pause

			if !m.pause {
				return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))
			}
			return m, nil
		}

		if m.win == nil && m.lose == nil && !m.resizeWarning && !m.pause {
			for i, binding := range m.keys.Letters {
				if !key.Matches(msg, binding) {
					continue
				}
				letter := rune('a' + i)

				if slices.Contains(m.gophersFirstChar, letter) && m.selected == nil {
					selected := &m.gophers[slices.Index(m.gophersFirstChar, letter)]
					if selected.Alive {
						m.selected = selected
						m.selected.DisplayWord = m.selected.DisplayWord[1:]
						return m, nil
					}
				}
				if m.selected != nil && len(m.selected.DisplayWord) > 0 && m.selected.DisplayWordRunes()[0] == letter {
					m.selected.DisplayWord = m.selected.DisplayWord[1:]
					return m, nil
				}
				break
			}
		}
	}

	m.initGophers()

	return m, nil
}

func (m *model) clearGophers() {
	m.gophers = make([]gopher, 0)
	m.gophersFirstChar = make([]rune, 0)
	m.selected = nil
}

func (m *model) initGophers() {
	gopherCount := 10

	if m.height > 0 && len(m.gophers) == 0 {
		segmentWidth := m.width / gopherCount
		words := m.pickUniqueWords(gopherCount)
		for i := 0; i < gopherCount; i++ {
			segmentStart := i * segmentWidth                                       // left edge of this gopher's segment
			segmentMargin := 1                                                     // columns reserved on each side to prevent adjacency
			usableWidth := segmentWidth - (segmentMargin * 2)                      // placeable range within the segment after margins
			gopherSpacing := segmentStart + segmentMargin + rand.IntN(usableWidth) // final X: start + margin + random offset
			m.gophers = append(m.gophers, gopher{X: gopherSpacing, Y: (m.height - m.topPadding), Word: words[i], DisplayWord: words[i], Alive: true})
			m.gophersFirstChar = append(m.gophersFirstChar, []rune(words[i])[0])
		}
	}
}

func (m *model) pickUniqueWords(n int) []string {
	var wordList []string
	if m.wave <= 3 {
		wordList = append(wordList, easyWordList...)
	}

	if m.wave > 3 && m.wave <= 7 {
		wordList = append(wordList, mediumWordList...)
	}

	if m.wave > 7 {
		wordList = append(wordList, hardWordList...)
	}

	words := make([]string, 0, n)
	usedLetters := map[byte]bool{}
	shuffled := rand.Perm(len(wordList))
	for _, idx := range shuffled {
		w := wordList[idx]
		if len(w) == 0 || usedLetters[w[0]] {
			continue
		}
		usedLetters[w[0]] = true
		words = append(words, w)
		if len(words) == n {
			break
		}
	}
	return words
}
