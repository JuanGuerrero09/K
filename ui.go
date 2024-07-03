package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	// "github.com/charmbracelet/huh"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

type Question struct {
	question string
	answer   string
	input    Input
}

func NewQuestion(question string) Question {
	return Question{question: question}
}

func NewShortQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewShortAnswerField()
	q.input = field
	return q
}

func NewLongQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewLongAnswerField()
	q.input = field
	return q
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

type Main struct {
	index     int
	questions []Question
	height    int
	width     int
	styles    *Styles
	done      bool
}

func NewMain(questions []Question) *Main {
	styles := DefaultStyles()
	answerField := textinput.New()
	answerField.Focus()
	answerField.Placeholder = "Your answer here: "
	return &Main{questions: questions, styles: styles}
}

func (m Main) Init() tea.Cmd {
	return nil
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			}
			current.answer = current.input.Value()
			log.Printf("question: %s, answer: %s", current.question, current.answer)
			m.Next()
			return m, current.input.Blur
		}
	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}

func (m *Main) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (m Main) View() string {
	if m.done {
		var output string
		for _, q := range m.questions {
			output += fmt.Sprintf("%s, %s\n", q.question, q.answer)
		}
		return output
	}
	if m.width == 0 {
		return "loading..."
	}
	currentQuestion := m.questions[m.index]
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			currentQuestion.question,
			m.styles.InputField.Render(currentQuestion.input.View())),
	)
}

func teaui() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := os.Getenv("NAME")
	partnerName := os.Getenv("PARTNERNAME")

	log.Printf("The name is: %s", name)
	log.Printf("The partnername is: %s", partnerName)

	questions := []Question{NewShortQuestion("What's your name?"), NewShortQuestion("What's your partner name?"), NewLongQuestion("What do you love about him?")}
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
