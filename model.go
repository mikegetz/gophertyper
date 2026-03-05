package main

import (
	_ "embed"
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
	gophers []gopher

	// terminal dimensions
	width      int
	height     int
	topPadding int

	keys keyMap
}

type gopherType int

const (
	word gopherType = iota
	gopherIcon
	gopherPath
)

type gopher struct {
	X, Y int
	Word string
	Type gopherType
}

type keyMap struct {
	Quit key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c/esc", "quit"),
	),
}

func initialModel() model {
	model := model{
		keys:       keys,
		topPadding: 8,
	}

	return model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(moveGophers(time.Millisecond * 500))
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
