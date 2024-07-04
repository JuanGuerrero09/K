package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ToDo struct {
	isDone         bool
	Text           string
	Category       string
	CompletionDate time.Time
}

func main() {
	file, err := os.Open("list.txt")
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

	// for _, todo := range todos {
	// 	fmt.Printf("%+v\n", todo)
	// }
}

func ParseFile(file *os.File) ([]ToDo, error) {
	var todos []ToDo
	var currentCategory string
	startParsing := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !startParsing {
			if strings.HasPrefix(line, "###") {
				startParsing = true
			}
			continue
		}
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "@") {
			currentCategory = strings.TrimPrefix(line, "@")
		} else if strings.HasPrefix(line, "<E>") {
			currentCategory = "Bonus España"
		} else if strings.HasPrefix(line, "<EU>") {
			currentCategory = "Bonus Europa"
		} else if strings.HasPrefix(line, "-") {
			todo := ToDo{
				isDone:   false,
				Text:     strings.TrimSpace(strings.TrimPrefix(line, "-")),
				Category: currentCategory,
			}
			todos = append(todos, todo)
		} else if strings.HasPrefix(line, "F") {
			parts := strings.Split(line, "/")
			if len(parts) > 2 {
				return nil, fmt.Errorf("invalid format for completed task: %s", line)
			}
			dateStr := strings.TrimSpace(parts[1])
			completationDate, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format: %s", dateStr)
			}
			todo := ToDo{
				isDone:         true,
				Text:           strings.TrimSpace(strings.TrimPrefix(parts[0], "F")),
				Category:       currentCategory,
				CompletionDate: completationDate,
			}
			todos = append(todos, todo)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}

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
	
	for i := start; i < end; i++ {
		todo := m.todoList[i]
		cursor := " "
		checked := " "
		date := ""
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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.cursor < len(m.todoList)-1 {
				m.cursor++
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
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
