package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/huh"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

type Main struct {
	index       int
	questions   []string
	height      int
	width       int
	answerField textinput.Model
	styles      *Styles
}

func NewMain(questions []string) *Main {
	styles := DefaultStyles()
	answerField := textinput.New()
	answerField.Focus()
	answerField.Placeholder = "Your answer here: "
	return &Main{questions: questions, answerField: answerField, styles: styles}
}

func (m Main) Init() tea.Cmd {
	return nil
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.index++
			m.answerField.SetValue("done!")
			return m, nil
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m Main) View() string {
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index], 
			m.styles.InputField.Render(m.answerField.View())),
	)
}

func main() {

	questions := []string{"hi", "hix2"}
	m := NewMain(questions)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal: ", err)
		os.Exit(0)
	}
	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())

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
