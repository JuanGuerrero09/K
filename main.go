package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ToDo struct {
	isDone         bool
	Text           string
	Category       string
	Points         int
	CompletionDate time.Time
}

func main() {
	key := "thisis32bitlongpassphraseimusing"
	// To be deleted when final txt is done
	todos, f := getEncryptedTodos(key)

	listModel := initialListmodel(todos)

	p := tea.NewProgram(listModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	encryptTodos(key, todos, f)

}
