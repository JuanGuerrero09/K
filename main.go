package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/huh"
)

type Main struct {
}

func (m Main) Init() tea.Cmd {
	return nil
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Main) View() string {
	return "Hello Kath"
}

func main() {

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal: ", err)
		os.Exit(0)
	}
	defer f.Close()

	p := tea.NewProgram(Main{}, tea.WithAltScreen())

	p.Run()

	// var name string
	// var name2 string
	// form := huh.NewForm(
	// 	huh.NewGroup(huh.NewInput().Description("What should we call you?").Value(&name), huh.NewInput().Description("What should we call you?").Value(&name2)),
	// )

	// err := form.Run()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Println("Welcome, " + name + "!")
	// fmt.Println("Welcome, " + name2 + "!")
}
