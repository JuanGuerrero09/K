package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ToDo struct {
	isDone         bool
	isHidden       bool
	Text           string
	Category       string
	Points         int
	CompletionDate time.Time
}

func main() {
	listModel := initialListmodel()

	p := tea.NewProgram(listModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
