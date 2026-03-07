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
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Letters[0]): // a
			if slices.Contains(m.gophersFirstChar, 'a') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'a')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'a' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}

		case key.Matches(msg, m.keys.Letters[1]): // b
			if slices.Contains(m.gophersFirstChar, 'b') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'b')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'b' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[2]): // c
			if slices.Contains(m.gophersFirstChar, 'c') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'c')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'c' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[3]): // d
			if slices.Contains(m.gophersFirstChar, 'd') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'd')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'd' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[4]): // e
			if slices.Contains(m.gophersFirstChar, 'e') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'e')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'e' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[5]): // f
			if slices.Contains(m.gophersFirstChar, 'f') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'f')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'f' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[6]): // g
			if slices.Contains(m.gophersFirstChar, 'g') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'g')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'g' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[7]): // h
			if slices.Contains(m.gophersFirstChar, 'h') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'h')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'h' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[8]): // i
			if slices.Contains(m.gophersFirstChar, 'i') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'i')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'i' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[9]): // j
			if slices.Contains(m.gophersFirstChar, 'j') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'j')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'j' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[10]): // k
			if slices.Contains(m.gophersFirstChar, 'k') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'k')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'k' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[11]): // l
			if slices.Contains(m.gophersFirstChar, 'l') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'l')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'l' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[12]): // m
			if slices.Contains(m.gophersFirstChar, 'm') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'm')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'm' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[13]): // n
			if slices.Contains(m.gophersFirstChar, 'n') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'n')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'n' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[14]): // o
			if slices.Contains(m.gophersFirstChar, 'o') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'o')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'o' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[15]): // p
			if slices.Contains(m.gophersFirstChar, 'p') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'p')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'p' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[16]): // q
			if slices.Contains(m.gophersFirstChar, 'q') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'q')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'q' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[17]): // r
			if slices.Contains(m.gophersFirstChar, 'r') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'r')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'r' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[18]): // s
			if slices.Contains(m.gophersFirstChar, 's') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 's')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 's' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[19]): // t
			if slices.Contains(m.gophersFirstChar, 't') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 't')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 't' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[20]): // u
			if slices.Contains(m.gophersFirstChar, 'u') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'u')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'u' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[21]): // v
			if slices.Contains(m.gophersFirstChar, 'v') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'v')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'v' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[22]): // w
			if slices.Contains(m.gophersFirstChar, 'w') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'w')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'w' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[23]): // x
			if slices.Contains(m.gophersFirstChar, 'x') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'x')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'x' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[24]): // y
			if slices.Contains(m.gophersFirstChar, 'y') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'y')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'y' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
		case key.Matches(msg, m.keys.Letters[25]): // z
			if slices.Contains(m.gophersFirstChar, 'z') && m.selected == nil {
				m.selected = &m.gophers[slices.Index(m.gophersFirstChar, 'z')]
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
			}
			if m.selected != nil && m.selected.DisplayWordRunes()[0] == 'z' {
				m.selected.DisplayWord = m.selected.DisplayWord[1:]
				return m, nil
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
