package cmd


import (
	"fmt"
	"os"

	"console-app/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func Execute() {
	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}