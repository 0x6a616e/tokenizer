package internal

import (
	"fmt"
	"strings"

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
	choices    []string
	cursor     int
	phase      Phase
	textarea   textarea.Model
	err        error
	tokenizers []Tokenizer
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
			m.phase = ShowingResults
			for _, line := range strings.Split(m.textarea.Value(), "\n") {
				tokenizer := NewTokenizer()
				tokenizer.Tokenize(line)
				m.tokenizers = append(m.tokenizers, tokenizer)
			}
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

func (m TeaModel) UpdateResults(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.phase {
	case ReadingOptions:
		return m.UpdateOptions(msg)
	case ReadingStd:
		return m.UpdateStd(msg)
	case ShowingResults:
		return m.UpdateResults(msg)
	}
	return m, nil
}

func (m TeaModel) ViewOptions() string {
	s := "Tokenizer :3\n\n"
	s += "Información de uso\n"
	s += "- Este tokenizer esta hecho para el lenguaje XXX\n"
	s += "- La entrada se procesa línea por línea\n"
	s += "\n¿Qué entrada desea usar?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPresiona q para salir.\n"

	return s
}

func (m TeaModel) ViewStd() string {
	s := "Ingresa el código a tokenizar\n\n"
	s += m.textarea.View()
	s += "\n\n(Ctrl+C para terminar)\n"

	return s
}

func (m TeaModel) ViewResults() string {
	s := "Resultados:\n\n"

	for _, tokenizer := range m.tokenizers {
		s += fmt.Sprintf("%v\n", tokenizer)
	}

	return s
}

func (m TeaModel) View() string {
	switch m.phase {
	case ReadingOptions:
		return m.ViewOptions()
	case ReadingStd:
		return m.ViewStd()
	case ShowingResults:
		return m.ViewResults()
	}
	return ""
}

func NewModel() TeaModel {
	ti := textarea.New()
	ti.Placeholder = "Código a procesar"
	ti.SetHeight(10)
	ti.SetWidth(80)

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
