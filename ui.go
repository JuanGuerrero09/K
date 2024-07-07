package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type listScreenModel struct {
	todoList []ToDo
	cursor   int
	viewport viewport.Model
	width, height int
}

func initialListmodel(todos []ToDo) *listScreenModel {
	const width = 78

	vp := viewport.New(width, 25)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	return &listScreenModel{
		todoList: todos,
		cursor:   0,
		viewport: vp,
	}
}
func (m listScreenModel) Init() tea.Cmd {
	return nil
}

func (m listScreenModel) View() string {
	l := "List of items: \n"
	start := 0
	end := m.viewport.Height - 4
	
	if m.cursor >= end {
		start = m.cursor - (m.viewport.Height - 4) + 1
		end = m.cursor + 1
	}

	end = min(end, len(m.todoList))
	var currentCategory string
	
	for i := start; i < end; i++ {
		todo := m.todoList[i]
		cursor := " "
		checked := " "
		date := ""
		if currentCategory == "" {
			currentCategory = todo.Category
			l += fmt.Sprintf("%s \n", currentCategory)
		}
		if currentCategory != todo.Category {
			currentCategory = todo.Category
			l += fmt.Sprintf("%s \n", currentCategory)
		}
		if todo.isDone {
			date = todo.CompletionDate.Format("2006-01-02")
			checked = "x"
		}
		if m.cursor == i {
			cursor = ">"
		}
		l += fmt.Sprintf("%s [%s] %s %s\n", cursor, checked, todo.Text, date)
	}


 	l += "\nPress q to quit.\n"

	m.viewport.SetContent(l)
	return lipgloss.NewStyle().Width(m.width).AlignHorizontal(lipgloss.Center).Render(m.viewport.View())
}

func (m listScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = m.width / 2
		if m.width < 78 {
				m.viewport.Width = m.width - 2
		}
		m.viewport.Height = m.height - 2
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q" :
			return m, tea.Quit
		case "down", "j", "s":
			if m.cursor < len(m.todoList)-1 {
				m.cursor++
			}
		case "up", "k", "w":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter", " ":
			if m.cursor < len(m.todoList) {
					m.todoList[m.cursor].isDone = !m.todoList[m.cursor].isDone
					m.todoList[m.cursor].CompletionDate = time.Now()
			}
		default:
			return m, cmd
		}
		
	}
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
