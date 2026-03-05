package main

import (
	"math/rand"
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
		randomGopher := rand.Intn(len(m.gophers))
		timeMultiplier := 50 - m.wave

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
		case key.Matches(msg, m.keys.Letters[0]): // a
			if slices.Contains(m.gophersFirstChar, 'a') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'a')]
			}

		case key.Matches(msg, m.keys.Letters[1]): // b
			if slices.Contains(m.gophersFirstChar, 'b') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'b')]
			}
		case key.Matches(msg, m.keys.Letters[2]): // c
			if slices.Contains(m.gophersFirstChar, 'c') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'c')]
			}
		case key.Matches(msg, m.keys.Letters[3]): // d
			if slices.Contains(m.gophersFirstChar, 'd') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'd')]
			}
		case key.Matches(msg, m.keys.Letters[4]): // e
			if slices.Contains(m.gophersFirstChar, 'e') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'e')]
			}
		case key.Matches(msg, m.keys.Letters[5]): // f
			if slices.Contains(m.gophersFirstChar, 'f') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'f')]
			}
		case key.Matches(msg, m.keys.Letters[6]): // g
			if slices.Contains(m.gophersFirstChar, 'g') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'g')]
			}
		case key.Matches(msg, m.keys.Letters[7]): // h
			if slices.Contains(m.gophersFirstChar, 'h') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'h')]
			}
		case key.Matches(msg, m.keys.Letters[8]): // i
			if slices.Contains(m.gophersFirstChar, 'i') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'i')]
			}
		case key.Matches(msg, m.keys.Letters[9]): // j
			if slices.Contains(m.gophersFirstChar, 'j') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'j')]
			}
		case key.Matches(msg, m.keys.Letters[10]): // k
			if slices.Contains(m.gophersFirstChar, 'k') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'k')]
			}
		case key.Matches(msg, m.keys.Letters[11]): // l
			if slices.Contains(m.gophersFirstChar, 'l') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'l')]
			}
		case key.Matches(msg, m.keys.Letters[12]): // m
			if slices.Contains(m.gophersFirstChar, 'm') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'm')]
			}
		case key.Matches(msg, m.keys.Letters[13]): // n
			if slices.Contains(m.gophersFirstChar, 'n') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'n')]
			}
		case key.Matches(msg, m.keys.Letters[14]): // o
			if slices.Contains(m.gophersFirstChar, 'o') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'o')]
			}
		case key.Matches(msg, m.keys.Letters[15]): // p
			if slices.Contains(m.gophersFirstChar, 'p') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'p')]
			}
		case key.Matches(msg, m.keys.Letters[16]): // q
			if slices.Contains(m.gophersFirstChar, 'q') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'q')]
			}
		case key.Matches(msg, m.keys.Letters[17]): // r
			if slices.Contains(m.gophersFirstChar, 'r') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'r')]
			}
		case key.Matches(msg, m.keys.Letters[18]): // s
			if slices.Contains(m.gophersFirstChar, 's') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 's')]
			}
		case key.Matches(msg, m.keys.Letters[19]): // t
			if slices.Contains(m.gophersFirstChar, 't') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 't')]
			}
		case key.Matches(msg, m.keys.Letters[20]): // u
			if slices.Contains(m.gophersFirstChar, 'u') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'u')]
			}
		case key.Matches(msg, m.keys.Letters[21]): // v
			if slices.Contains(m.gophersFirstChar, 'v') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'v')]
			}
		case key.Matches(msg, m.keys.Letters[22]): // w
			if slices.Contains(m.gophersFirstChar, 'w') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'w')]
			}
		case key.Matches(msg, m.keys.Letters[23]): // x
			if slices.Contains(m.gophersFirstChar, 'x') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'x')]
			}
		case key.Matches(msg, m.keys.Letters[24]): // y
			if slices.Contains(m.gophersFirstChar, 'y') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'y')]
			}
		case key.Matches(msg, m.keys.Letters[25]): // z
			if slices.Contains(m.gophersFirstChar, 'z') {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'z')]
			}
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
			gopherSpacing := segmentStart + segmentMargin + rand.Intn(usableWidth) // final X: start + margin + random offset
			m.gophers = append(m.gophers, gopher{X: gopherSpacing, Y: (m.height - m.topPadding), Word: words[i]})
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
