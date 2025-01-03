package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gocyclops/cyclops/internal/generator"
	"github.com/gocyclops/cyclops/utils"
)

type errMsg struct {
	err error
}

func generateProject(project generator.Project) tea.Cmd {
	return func() tea.Msg {
		err := project.Generate()
    return errMsg{err: err}
	}
}

func (m Model) updateInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	cleanProjectName := utils.CleanProjectName(m.projectInput.Value())
	cleanRepoURL := utils.CleanRepoUrl(m.repoURL.Value())

	switch msg.String() {
	case "tab", "shift+tab", "up", "down", "enter":
		if msg.String() == "enter" && m.activeInput == 1 {
			if valid, errMsg := utils.ValidateInputs(cleanProjectName, cleanRepoURL); !valid {
				m.err = fmt.Errorf("%s", errMsg)
				return m, nil
			}
		
			m.state = "framework"
			return m, nil
		}

		if m.activeInput == 0 {
			m.projectInput.Blur()
      m.repoURL.Focus()
      m.activeInput = 1
		} else {
			m.repoURL.Blur()
      m.projectInput.Focus()
      m.activeInput = 0
		}

		return m, nil
	default:
		var cmd tea.Cmd
		if m.activeInput == 0 {
			m.projectInput, cmd = m.projectInput.Update(msg)
	} else {
			m.repoURL, cmd = m.repoURL.Update(msg)
	}
		return m, cmd
	}
}

func (m Model) updateFramework(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0{
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.frameworks)-1 {
			m.cursor++
		}
	case "enter":
		m.framework = m.frameworks[m.cursor]
		m.state = "features"
		m.cursor = 0
	case "esc":
		m.state = "input"
	}
	return m, nil
}

func (m Model) updateFeatures(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.features)-1 {
			m.cursor++
		}
	case "enter":
		feature := m.features[m.cursor]
		if m.selected[feature] {
			delete(m.selected, feature)
		} else {
			m.selected[feature] = true
		}
	case "tab":
		m.state = "confirm"
	case "esc":
		m.state = "framework"
	}
	return m, nil
}

func (m *Model) updateConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if !m.generating {
			project := generator.Project{
				Name:      m.projectInput.Value(),
				Framework: m.framework,
				Features:  m.selected,
			}
			
			m.generating = true
			m.progress = "Generating project..."
			return m, generateProject(project)	
		}
	case "esc":
		if !m.generating {
			m.state = "features"
		}
	}
	return m, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
    if msg.String() == "ctrl+c" {
      return m, tea.Quit
    }
    
		switch m.state {
			case "input":
				return m.updateInput(msg)
			case "framework":
				return m.updateFramework(msg)
			case "features":
				return m.updateFeatures(msg)
			case "confirm":
				return m.updateConfirm(msg)
      case "done":
        if msg.String() == "enter" {
					return m, tea.Quit
			}
		}

	case errMsg:
		if msg.err != nil {
			m.err = msg.err
			m.progress = fmt.Sprintf("Error: %v", msg.err)
		}

		m.progress = "Project generation complete! 🎉\n\nPress Enter to exit."
		m.generating = false
		m.state = "done"
		return m, nil
	}

	return m, nil
}
