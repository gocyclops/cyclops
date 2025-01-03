package create

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gocyclops/cyclops/internal/tui"
)

func Execute() error {
	program := tea.NewProgram(tui.NewModel())
	_, err := program.Run()
	return err
}
