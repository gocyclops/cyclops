package tui

import (
	"fmt"
)

func (m Model) viewInput() string {
	return fmt.Sprintf(
		"Enter project name:\n\n%s\n\nPress Enter to continue.",
		m.projectInput.View(),
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
	return s + "\nUse arrow keys to navigate. Press Enter to select."
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
	return s + "\nUse arrow keys to navigate, Enter to toggle, Tab to confirm."
}

func (m Model) viewConfirm() string {
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
	return "Project generation complete! 🎉\n"
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
