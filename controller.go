package main

import (
	"math/rand/v2"
	"slices"
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
		randomGopher := m.RandomLivingGopher()
		timeMultiplier := 50 - m.wave

		if randomGopher == nil {
			return m, nil
		}

		if randomGopher.Y > 0 {
			randomGopher.Y--
		} else {
			m.lose = randomGopher
			return m, nil
		}

		if m.selected != nil {
			if len(m.selected.DisplayWord) == 0 {
				m.selected.Alive = false
				m.selected = nil
			}
		}

		return m, moveGophers(time.Millisecond * time.Duration(timeMultiplier))

	case tea.KeyPressMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

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
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == letter {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			break
		}
	}

	m.initGophers()

	return m, nil
}

func (m *model) initGophers() {
	gopherCount := 10

	if m.height > 0 && len(m.gophers) == 0 {
		segmentWidth := m.width / gopherCount
		words := pickUniqueWords(easyWordList, gopherCount)
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

func pickUniqueWords(wordList []string, n int) []string {
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
