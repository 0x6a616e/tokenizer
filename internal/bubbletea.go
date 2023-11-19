package internal

import (
	"math"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("15"))

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
	phase     Phase
	textarea  textarea.Model
	err       error
	tokenizer Tokenizer
	table     table.Model
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

func (m TeaModel) UpdateWelcome(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.phase = ReadingInput
			m.textarea.Focus()
			return m, textarea.Blink
		}
	}
	return m, nil
}

func (m *TeaModel) MakeTable() {
	c := []table.Column{
		{Title: "", Width: 0},
		{Title: "Tipo", Width: 4},
		{Title: "Contenido", Width: 9},
	}
	r := []table.Row{}
	for _, t := range m.tokenizer.Tokens {
		var row table.Row
		switch t.Type {
		case Undefined:
			row = table.Row{"X", "Indefinido", t.Content}
		case Identifier:
			row = table.Row{"O", "Identificador", t.Content}
		case Keyword:
			row = table.Row{"O", "Reservada", t.Content}
		case Operator:
			row = table.Row{"O", "Operador", t.Content}
		case StringLiteral:
			row = table.Row{"O", "Cadena literal", t.Content}
		case IntegerLiteral:
			row = table.Row{"O", "Entero literal", t.Content}
		case DecimalLiteral:
			row = table.Row{"O", "Decimal literal", t.Content}
		case Invalid:
			row = table.Row{"X", "Inválido", t.Content}
		}
		r = append(r, row)
		c[0].Width = int(math.Max(float64(c[0].Width), float64(len(row[0]))))
		c[1].Width = int(math.Max(float64(c[1].Width), float64(len(row[1]))))
		c[2].Width = int(math.Max(float64(c[2].Width), float64(len(row[2]))))
	}
	m.table.SetColumns(c)
	m.table.SetRows(r)
	m.table.SetHeight(int(math.Min(float64(len(r)), 10)))
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
			m.tokenizer.Tokenize(m.textarea.Value())
			m.MakeTable()
			m.table.Focus()
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
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.phase = ShowingWelcome
			m.textarea.SetValue("")
			m.err = nil
			m.tokenizer = Tokenizer{}
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
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
	s += "- Este tokenizer esta hecho para el lenguaje Go\n"
	s += "- Solo reconoce un subconjunto de los tokens que Go puede reconocer\n"
	s += "- Cada token se separa por un espacio ( ), un punto y coma (;) o un salto de línea\n"
	s += "\n"
	s += "(Enter para iniciar)\n"
	s += "(Ctrl+C para salir)\n"

	return s
}

func (m TeaModel) ViewInput() string {
	s := "Ingresa el código a tokenizar\n"
	s += "\n"
	s += m.textarea.View() + "\n"
	s += "\n"
	s += "(Ctrl+C para terminar de leer)\n"

	return s
}

func (m TeaModel) ViewResults() string {
	s := "Tokens:\n"
	s += "\n"
	s += baseStyle.Render(m.table.View()) + "\n"
	s += "\n"
	s += "(Flecha arriba/abajo para desplazarse)\n"
	s += "(Enter para volver al inicio)\n"
	s += "(Ctrl+C para terminar el programa)\n"

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

	t := table.New()

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("15")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("15")).
		Bold(false)
	t.SetStyles(s)

	return TeaModel{
		phase:    ShowingWelcome,
		textarea: ti,
		err:      nil,
		table:    t,
	}
}
