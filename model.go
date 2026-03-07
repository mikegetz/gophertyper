package main

import (
	_ "embed"
	"math/rand/v2"
	"strings"
	"time"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

//go:embed words/easy.txt
var easyWords string

//go:embed words/medium.txt
var mediumWords string

//go:embed words/hard.txt
var hardWords string

var (
	easyWordList   = strings.Split(easyWords, "\n")
	mediumWordList = strings.Split(mediumWords, "\n")
	hardWordList   = strings.Split(hardWords, "\n")
)

type model struct {
	gophers          []gopher
	gophersFirstChar []rune
	wave             int
	lose             *gopher
	win              *gopher
	selected         *gopher

	// terminal dimensions
	width      int
	height     int
	topPadding int

	keys keyMap
}

func initialModel() model {
	model := model{
		keys:       keys,
		topPadding: 8,
		lose:       nil,
		win:        nil,
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

type clockTickMsg struct{}

func clockTick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return clockTickMsg{}
	})
}

func moveGophers(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
