package internal

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	Phase  int
	errMsg error
)

const (
	ReadingOptions Phase = iota
	ReadingStd
	ReadingFile
	ShowingResults
)

type TeaModel struct {
	choices  []string
	cursor   int
	phase    Phase
	textarea textarea.Model
	err      error
}

func (m TeaModel) Init() tea.Cmd {
	return nil
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
			switch m.phase {
			case ReadingStd:
				m.textarea.Focus()
				return m, textarea.Blink
			}
		}
	}
	return m, nil
}

func (m TeaModel) UpdateStd(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.phase {
	case ReadingOptions:
		return m.UpdateOptions(msg)
	case ReadingStd:
		return m.UpdateStd(msg)
	}
	return m, nil
}

// TODO: Desplegar mensaje de bienvenida y decir que se procesa por línea
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

func (m TeaModel) ViewStd() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func (m TeaModel) View() string {
	switch m.phase {
	case ReadingOptions:
		return m.ViewOptions()
	case ReadingStd:
		return m.ViewStd()
	}
	return ""
}

func NewModel() TeaModel {
	ti := textarea.New()
	ti.Placeholder = "Ingresa el código a procesar"

	return TeaModel{
		choices: []string{
			"Leer desde consola",
			"Leer desde un archivo",
		},
		phase:    ReadingOptions,
		textarea: ti,
		err:      nil,
	}
}
