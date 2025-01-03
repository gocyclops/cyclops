package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	projectInput textinput.Model
	repoURL      textinput.Model
	activeInput  int
	frameworks   []string
	features     []string
	cursor       int
	selected     map[string]bool
	state        string
	projectName  string
	framework    string
	err          error
	generating   bool
	progress     string
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

	link := textinput.New()
	link.Placeholder = "github.com/username/repo"

	return Model{
		projectInput: ti,
		repoURL:      link,
		activeInput:  0,
		frameworks:   []string{"fiber", "gin", "chi", "gorilla/mux"},
		features:     []string{"redis", "s3", "mail", "auth"},
		selected:     make(map[string]bool),
		state:        StateInput,
	}
}
