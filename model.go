package main

import (
	_ "embed"
	"math/rand/v2"
	"slices"
	"strings"
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

const minTerminalWidth = 100

//go:embed words/easy.txt
var easyWords string

//go:embed words/medium.txt
var mediumWords string

//go:embed words/hard.txt
var hardWords string

//go:embed words/win
var winText string

//go:embed words/wave
var wave string

//go:embed words/0
var waveZero string

//go:embed words/1
var waveOne string

//go:embed words/2
var waveTwo string

//go:embed words/3
var waveThree string

//go:embed words/4
var waveFour string

//go:embed words/5
var waveFive string

//go:embed words/6
var waveSix string

//go:embed words/7
var waveSeven string

//go:embed words/8
var waveEight string

//go:embed words/9
var waveNine string

var waveNumbers = []string{waveZero, waveOne, waveTwo, waveThree, waveFour, waveFive, waveSix, waveSeven, waveEight, waveNine}

var (
	easyWordList   = strings.Split(easyWords, "\n")
	mediumWordList = strings.Split(mediumWords, "\n")
	hardWordList   = strings.Split(hardWords, "\n")
)

type model struct {

	// game state
	gophers            []gopher
	gophersFirstChar   []rune
	usedWords          []string
	wave               int
	waveTransition     bool
	winTransition      bool
	loseTransition     bool
	timeMultiplier     int
	userTimeMultiplier int
	pause              bool
	lose               *gopher
	win                *gopher
	selected           *gopher

	// terminal dimensions
	width         int
	height        int
	topPadding    int
	resizeWarning bool
	isDark        bool

	// stats
	correctKeypresses int
	keypresses        int
	killCount         int
	gpm               string
	wpm               string
	gpmStart          time.Time
	gpmEnd            time.Time
	pauseStart        time.Time
	pauseEnd          time.Time
	pauseDuration     time.Duration
	accuracy          string

	// key bindings
	help help.Model
	keys keyMap
}

func initialModel() model {
	h := help.New()

	h.ShowAll = false
	h.Styles = help.DefaultDarkStyles()

	model := model{
		help:              h,
		keys:              keys,
		topPadding:        10,
		lose:              nil,
		win:               nil,
		waveTransition:    false,
		winTransition:     false,
		wave:              0,
		killCount:         0,
		correctKeypresses: 0,
		gpm:               "0",
		wpm:               "0",
		accuracy:          "0%",
		keypresses:        0,
	}

	return model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(moveGophers(time.Millisecond * 500))
}

func (m model) randomLivingGopher() *gopher {
	var living []*gopher
	for i := range m.gophers {
		if m.gophers[i].Alive {
			living = append(living, &m.gophers[i])
		}
	}
	if len(living) == 0 {
		return nil
	}
	return living[rand.IntN(len(living))]
}

func (m model) livingGopherCount() int {
	count := 0
	for _, gopher := range m.gophers {
		if gopher.Alive {
			count++
		}
	}
	return count
}

type gopherType int

const (
	gopherIcon gopherType = iota
	gopherPath
)

type gopher struct {
	X, Y        int
	Word        string
	DisplayWord string
	Type        gopherType
	Alive       bool
}

func (g gopher) WordRunes() []rune {
	return []rune(g.Word)
}

func (g gopher) DisplayWordRunes() []rune {
	return []rune(g.DisplayWord)
}

type keyMap struct {
	Quit    key.Binding
	Pause   key.Binding
	Up      key.Binding
	Down    key.Binding
	Reset   key.Binding
	Letters [26]key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Pause, k.Up, k.Down, k.Reset}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Quit, k.Pause, k.Up, k.Down, k.Reset}}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c/esc", "quit"),
	),
	Pause: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("space", "pause"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "increase speed"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "decrease speed"),
	),
	Reset: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "reset"),
	),
	Letters: func() [26]key.Binding {
		var bindings [26]key.Binding
		for i := 0; i < 26; i++ {
			letter := string(rune('a' + i))
			bindings[i] = key.NewBinding(key.WithKeys(letter))
		}
		return bindings
	}(),
}

type tickMsg time.Time

type waveTransitionMsg struct{}

type winTransitionMsg struct{}

type loseTransitionMsg struct{}

func waveTransition(m *model, d time.Duration) tea.Cmd {
	m.waveTransition = true
	return tea.Tick(d, func(time.Time) tea.Msg {
		return waveTransitionMsg{}
	})
}

func winTransition(m *model, d time.Duration) tea.Cmd {
	m.winTransition = true
	return tea.Tick(d, func(time.Time) tea.Msg {
		return winTransitionMsg{}
	})
}

func loseTransition(m *model, d time.Duration) tea.Cmd {
	m.loseTransition = true
	return tea.Tick(d, func(time.Time) tea.Msg {
		return loseTransitionMsg{}
	})
}

func moveGophers(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *model) reset() {
	m.wave = 0
	m.userTimeMultiplier = 0
	m.correctKeypresses = 0
	m.keypresses = 0
	m.killCount = 0
	m.clearGophers()
	m.initGophers()
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

	if m.wave < 5 {
		wordList = append(wordList, easyWordList...)
	}

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

		alreadyUsed := slices.Contains(m.usedWords, w)
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
