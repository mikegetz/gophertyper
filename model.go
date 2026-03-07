package main

import (
	_ "embed"
	"math/rand/v2"
	"strings"
	"time"

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
	gophers          []gopher
	gophersFirstChar []rune
	wave             int
	waveTransition   bool
	winTransition    bool
	timeMultiplier   int
	lose             *gopher
	win              *gopher
	selected         *gopher

	// terminal dimensions
	width         int
	height        int
	topPadding    int
	resizeWarning bool

	keys keyMap
}

func initialModel() model {
	model := model{
		keys:           keys,
		topPadding:     10,
		lose:           nil,
		win:            nil,
		waveTransition: false,
		winTransition:  false,
		wave:           0,
	}

	return model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(moveGophers(time.Millisecond * 500))
}

func (m model) RandomLivingGopher() *gopher {
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

func (m model) LivingGopherCount() int {
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
	Letters [26]key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c/esc", "quit"),
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
	return tea.Tick(d, func(time.Time) tea.Msg {
		return loseTransitionMsg{}
	})
}

func moveGophers(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
