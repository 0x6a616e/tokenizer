package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/0x6a616e/tokenizer/internal"
)

func main() {
	p := tea.NewProgram(internal.NewModel())
	if _, err := p.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
