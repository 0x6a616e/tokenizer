package main

import (
	"log"
	"os"

	"github.com/0x6a616e/tokenizer/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(internal.NewModel())
	if _, err := p.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
