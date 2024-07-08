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
	file, err := os.Open("newfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	todos, err := ParseFile(file)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	listModel := initialListmodel(todos)

	p := tea.NewProgram(listModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	saveFile("newfile.txt", todos)

	// for _, todo := range todos {
	// 	fmt.Printf("%+v\n", todo)
	// }
}
