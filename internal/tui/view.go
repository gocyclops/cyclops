package tui

import (
	"fmt"
)

func getInputHint(activeInput int) string {
	if activeInput == 0 {
		return "Tab to switch to repository • Enter to submit project name"
	}
	return "Tab to switch to project name • Enter to continue"
}

func (m Model) viewInput() string {
	var errorMsg string
	if m.err != nil {
		errorMsg = "\n❌ " + m.err.Error()
	}

	return fmt.Sprintf(
		"Project name: (tab/shift+tab to switch)\n\n%s\n\nRemote repository:\n\n%s\n\n%s%s",
		m.projectInput.View(),
		m.repoURL.View(),
		getInputHint(m.activeInput),
		errorMsg,
	)
}

func (m Model) viewFramework() string {
	s := "Select a framework:\n\n"
	for i, fw := range m.frameworks {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, fw)
	}
	return s + "\nUse arrow keys to navigate, Esc to go back. Press Enter to select."
}

func (m Model) viewFeatures() string {
	s := "Select features:\n\n"
	for i, feature := range m.features {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		selected := " "
		if m.selected[feature] {
			selected = "✔"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, feature)
	}
	return s + "\nUse arrow keys to navigate, Esc to go back, Enter to toggle, Tab to confirm."
}

func (m Model) viewConfirm() string {
	if m.generating {
		return fmt.Sprintf("Progress: %s", m.progress)
	}

	s := fmt.Sprintf("Project Name: %s\n", m.projectName)
	s += fmt.Sprintf("Framework: %s\n", m.framework)
	s += "Features:\n"
	for feature, selected := range m.selected {
		if selected {
			s += fmt.Sprintf(" - %s\n", feature)
		}
	}
	return s + "\nPress Enter to generate, Esc to go back."
}

func (m Model) viewDone() string {
	return m.progress
}

func (m Model) View() string {
	switch m.state {
		case "input":
			return m.viewInput()
		case "framework":
			return m.viewFramework()
		case "features":
			return m.viewFeatures()
		case "confirm":
			return m.viewConfirm()
		case "done":
			return m.viewDone()
		default:
			return "Error: Unknown state"
	}
}
