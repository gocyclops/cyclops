package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	projectInput	textinput.Model
	frameworks		[]string
	features			[]string
	cursor				int
	selected			map[string]bool
	state					string
	projectName 	string
  framework   	string
  //err         	error
}

func (m Model) Init() tea.Cmd {
	return nil
}

const (
	StateInput     = "input"
	StateFramework = "framework"
	StateFeatures  = "features"
	StateConfirm   = "confirm"
	StateDone      = "done"
)

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "project-name"
	ti.Focus()

	return Model{
		projectInput: ti,
		frameworks: []string{"fiber", "gin", "chi"},
		features: []string{"redis", "s3", "mail"},
		selected: make(map[string]bool),
		state: StateInput,
	}
}
