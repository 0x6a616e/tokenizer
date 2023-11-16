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
	ShowingWelcome Phase = iota
	ReadingInput
	ShowingResults
)

type TeaModel struct {
	phase      Phase
	textarea   textarea.Model
	err        error
	tokenizers []Tokenizer
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

func (m TeaModel) UpdateWelcome(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			m.phase = ReadingInput
			m.textarea.Focus()
			return m, textarea.Blink
		}
	}
	return m, nil
}

func (m TeaModel) UpdateInput(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case ShowingWelcome:
		return m.UpdateWelcome(msg)
	case ReadingInput:
		return m.UpdateInput(msg)
	case ShowingResults:
		return m.UpdateResults(msg)
	}
	return m, nil
}

func (m TeaModel) ViewWelcome() string {
	s := "Tokenizer :3\n"
	s += "\n"
	s += "Información de uso\n"
	s += "- Este tokenizer esta hecho para el lenguaje XXX\n"
	s += "- La entrada se procesa línea por línea\n"
	s += "\n"
	s += "Presiona Enter para iniciar\n"
	s += "Presiona q para salir.\n"

	return s
}

func (m TeaModel) ViewInput() string {
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
	case ShowingWelcome:
		return m.ViewWelcome()
	case ReadingInput:
		return m.ViewInput()
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
		phase:    ShowingWelcome,
		textarea: ti,
		err:      nil,
	}
}
