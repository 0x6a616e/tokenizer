package internal

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Phase int

const (
	ReadingOptions Phase = iota
	ReadingStd
	ReadingFile
)

type TeaModel struct {
	choices []string
	cursor  int
	phase   Phase
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

func (m TeaModel) ViewOptions() string {
	s := "¿Qué entrada desea usar?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m TeaModel) View() string {
	switch m.phase {
	case ReadingOptions:
		return m.ViewOptions()
	}
	return ""
}

func (m TeaModel) UpdateOptions(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.phase = Phase(m.cursor) + 1
			return m, nil
		}
	}
	return m, nil
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.phase {
	case ReadingOptions:
		return m.UpdateOptions(msg)
	}
	return m, nil
}

func NewModel() TeaModel {
	return TeaModel{
		choices: []string{
			"Leer desde consola",
			"Leer desde un archivo",
		},
		phase: ReadingOptions,
	}
}
