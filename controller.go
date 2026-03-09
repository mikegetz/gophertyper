package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
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
		m.pauseEnd = time.Now()
		m.pauseDuration += m.pauseEnd.Sub(m.pauseStart)
		return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))

	case tickMsg:
		if m.resizeWarning || m.win != nil || m.lose != nil || m.pause {
			return m, nil
		}
		if m.gpmStart.IsZero() {
			m.gpmStart = time.Now()
		}

		if m.selected != nil {
			if len(m.selected.DisplayWord) == 0 {
				m.selected.Alive = false
				m.killCount++
				m.selected = nil
			}
		}

		randomGopher := m.RandomLivingGopher()

		if randomGopher == nil {
			if m.wave > 9 {
				m.win = m.selected
				m.gpmEnd = time.Now()
				m.gpm = calculateGPM(m.gpmStart, m.gpmEnd, m.pauseDuration, m.killCount)
				m.wpm = calculateWPM(m.gpmStart, m.gpmEnd, m.pauseDuration, m.correctKeypresses, m.killCount)
				return m, winTransition(&m, time.Second*5)
			}
			m.pauseStart = time.Now()
			return m, waveTransition(&m, time.Second*3)
		}

		if randomGopher.Y > 0 {
			randomGopher.Y--
		} else {
			m.lose = randomGopher
			m.gpmEnd = time.Now()
			m.gpm = calculateGPM(m.gpmStart, m.gpmEnd, m.pauseDuration, m.killCount)
			m.wpm = calculateWPM(m.gpmStart, m.gpmEnd, m.pauseDuration, m.correctKeypresses, m.killCount)
			return m, loseTransition(&m, time.Second*15)
		}

		return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))

	case tea.KeyPressMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, m.keys.Pause) {
			m.pause = !m.pause

			if !m.pause {
				m.pauseEnd = time.Now()
				m.pauseDuration += m.pauseEnd.Sub(m.pauseStart)
				return m, moveGophers(time.Millisecond * time.Duration(m.timeMultiplier))
			}
			m.pauseStart = time.Now()
			return m, nil
		}

		if m.win == nil && m.lose == nil && !m.resizeWarning && !m.pause {
			for i, binding := range m.keys.Letters {
				if !key.Matches(msg, binding) {
					continue
				}
				m.keypresses++
				letter := rune('a' + i)

				if slices.Contains(m.gophersFirstChar, letter) && m.selected == nil {
					selected := &m.gophers[slices.Index(m.gophersFirstChar, letter)]
					if selected.Alive {
						m.correctKeypresses++
						m.selected = selected
						m.selected.DisplayWord = m.selected.DisplayWord[1:]
						return m, nil
					}
				}
				if m.selected != nil && len(m.selected.DisplayWord) > 0 && m.selected.DisplayWordRunes()[0] == letter {
					m.correctKeypresses++
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
		m.usedWords = append(m.usedWords, words...)
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
	wordList = append(wordList, easyWordList...)

	if m.wave > 2 {
		wordList = append(wordList, mediumWordList...)
	}

	if m.wave > 6 {
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

		alreadyUsed := false
		for _, uw := range m.usedWords {
			if w == uw {
				alreadyUsed = true
				break
			}
		}
		if alreadyUsed {
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

func calculateGPM(start, end time.Time, pauseDuration time.Duration, kills int) string {
	minutes := (end.Sub(start) - pauseDuration).Minutes()
	if minutes == 0 {
		return "0"
	}
	gpm := float64(kills) / minutes
	return strconv.Itoa(int(gpm))
}

func calculateWPM(start, end time.Time, pauseDuration time.Duration, correctChars int, completedWords int) string {
	minutes := (end.Sub(start) - pauseDuration).Minutes()
	if minutes == 0 {
		return "0"
	}
	wpm := float64(correctChars+completedWords) / 5.0 / minutes
	return fmt.Sprintf("%.0f", wpm)
}
