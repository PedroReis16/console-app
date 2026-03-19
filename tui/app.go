package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor int
	items  []string
	msg    string
}

func New() tea.Model {
	return model{
		cursor: 0,
		items:  []string{"Status", "Estudos", "Configurações", "Logs", "Sair"},
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch k := msg.(type) {
	case tea.KeyMsg:
		switch k.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			choice := m.items[m.cursor]
			m.msg = fmt.Sprintf("Selecionado: %s (q para sair)", choice)
			// Se quiser sair ao selecionar "Sair", descomente:
			// if choice == "Sair" { return m, tea.Quit }
		}
	}

	return m, nil
}

func (m model) View() string {
	title := lipgloss.NewStyle().Bold(true).Underline(true).Render("console-app TUI")

	body := ""
	for i, item := range m.items {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		style := lipgloss.NewStyle()
		if i == m.cursor {
			style = style.Reverse(true)
		}

		body += style.Render(fmt.Sprintf("%s %s", cursor, item)) + "\n"
	}

	footer := lipgloss.NewStyle().Faint(true).Render("Use ↑/↓ para navegar, Enter para escolher, q para sair")

	msg := ""
	if m.msg != "" {
		msg = lipgloss.NewStyle().Italic(true).Render(m.msg) + "\n\n"
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s%s",
		title,
		msg,
		body,
		footer,
	)
}

