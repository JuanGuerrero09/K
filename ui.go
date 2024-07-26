package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	CHARSET = "⠿⠾⠽⠼⠻⠺⠹⠸⠷⠶⠵⠴⠳⠲⠱⠰⠯⠮⠭⠬⠫⠪⠩⠨⠧⠦⠥⠤⠣⠢⠡⠠⠟⠞⠝⠜⠛⠚⠙⠘⠗⠖⠕⠔⠓⠒⠑⠐⠏⠎⠍⠌⠋⠊⠉⠈⠇⠆⠅⠄⠃⠂⠁"
	totalPoints int
)

func getRandomString(length int) string {
	letters := []rune(CHARSET)
	b := make([]rune, length)
	for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type listScreenModel struct {
	todoList       []ToDo
	activeTodoList []ToDo
	cursor         int
	viewport       viewport.Model
	width, height  int
	tempFile       *os.File
	key            string
}

func initialListmodel() *listScreenModel {
	const width = 85

	vp := viewport.New(width, 25)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	return &listScreenModel{
		cursor:   0,
		viewport: vp,
	}
}

type TodoIO []ToDo

func startIO() tea.Msg {
	var todos []ToDo
	return TodoIO(todos)
}

func (m listScreenModel) Init() tea.Cmd {
	return startIO
}

func getPoints(todos []ToDo) int {
	count := 0
	for _, todo := range todos {
		if todo.isDone {
			count += todo.Points
		} else {
			continue
		}
	}
	return count
}

func (m listScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case TodoIO:
		m.key = "thisis32bitlongpassphraseimusing"
		// To be deleted when final txt is done
		todos, f := getEncryptedTodos(m.key)
		m.todoList = todos
		m.tempFile = f
		m.activeTodoList = m.todoList
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.width/2 < 85 {
			m.viewport.Width = 85

		} else {
			m.viewport.Width = m.width / 2

		}
		if m.width < 85 {
			m.viewport.Width = m.width - 2
		}
		m.viewport.Height = m.height - 2
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			encryptTodos(m.key, m.todoList, m.tempFile)
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
			if m.cursor < len(m.activeTodoList) {
				if m.activeTodoList[m.cursor].isHidden {
					return m, cmd
				}
				m.activeTodoList[m.cursor].isDone = !m.activeTodoList[m.cursor].isDone
				m.activeTodoList[m.cursor].CompletionDate = time.Now()
				idx := slices.IndexFunc(m.todoList, func(t ToDo) bool {return t.Text == m.activeTodoList[m.cursor].Text})
				m.todoList[idx] = m.activeTodoList[m.cursor]
				return m, cmd
			}
		case "f":
			var activeTodo []ToDo
			for _, todo := range(m.todoList){
				if todo.isDone {
					activeTodo = append(activeTodo, todo)
				}
			}
			m.cursor = 0
			m.activeTodoList = activeTodo
		case "r":
			var activeTodo []ToDo
			for _, todo := range(m.todoList){
				if !todo.isDone {
					activeTodo = append(activeTodo, todo)
				}
			}
			m.cursor = 0
			m.activeTodoList = activeTodo
		case "a":
			m.activeTodoList = m.todoList
		default:
			return m, cmd
		}
	}
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m listScreenModel) View() string {
	
	if totalPoints == 0 {
		for _, todo := range(m.todoList){
			totalPoints += todo.Points
		}
	}

	if getPoints(m.todoList) >= 50 {
		for i := range(m.todoList){
			if m.todoList[i].Category == "Bonus Madrid" {
				m.todoList[i].isHidden = false
			}
		}
	}

	if getPoints(m.todoList) >= 70 {
		for i := range(m.todoList){
			if m.todoList[i].Category == "Bonus España" {
				m.todoList[i].isHidden = false
			}
		}
	}

	if getPoints(m.todoList) >= 100 {
		for i := range(m.todoList){
			if m.todoList[i].Category == "Bonus Europa" {
				m.todoList[i].isHidden = false
			}
		}
	}
	
	l := "Points: " + strconv.Itoa(getPoints(m.todoList)) + "/" + strconv.Itoa(totalPoints)
	remainSpaces := m.viewport.Width - len(l) - 23
	l += strings.Repeat(" ", remainSpaces)
	l += "Points    Date \n"
	start := 0
	end := m.viewport.Height - 4

	if m.cursor >= end {
		start = m.cursor - (m.viewport.Height - 4) + 1
		end = m.cursor + 1
	}

	end = min(end, len(m.activeTodoList))
	var currentCategory string
	// var style = lipgloss.NewStyle().
	// 	Bold(true).
	// 	Foreground(lipgloss.Color("#FAFAFA")).
	// 	Background(lipgloss.Color("#7D56F4")).
	// 	Width(22)

	for i := start; i < end ; i++ {
		var todoStr string
		todo := m.activeTodoList[i]
		remainSpaces := m.viewport.Width - utf8.RuneCountInString(todo.Text) - 26
		if remainSpaces < 0 {
			remainSpaces = 0
		}
		spacesStr := strings.Repeat(" ", remainSpaces)
		cursor := " "
		checked := " "
		date := ""
		categoryStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).Bold(true)
		completedTodoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#68DA37"))
		if currentCategory == "" {
			currentCategory = todo.Category
			categorySpaces := strings.Repeat(" ", m.viewport.Width - utf8.RuneCountInString(currentCategory) - 25)
			var currentCategoryPoints int
			var totalCategoryPoints int
			for _, todo := range(m.todoList){
				if todo.Category == currentCategory {
					totalCategoryPoints += todo.Points
					if todo.isDone {
						currentCategoryPoints += todo.Points
					}
				}
			} 
			c := fmt.Sprintf("%s %s %d / %d", currentCategory, categorySpaces, currentCategoryPoints, totalCategoryPoints)
			l += categoryStyle.Render(c)
			l += "\n"
		}
		if currentCategory != todo.Category {
			currentCategory = todo.Category
			categorySpaces := strings.Repeat(" ", m.viewport.Width - utf8.RuneCountInString(currentCategory) - 25)
			var currentCategoryPoints int
			var totalCategoryPoints int
			for _, todo := range(m.todoList){
				if todo.Category == currentCategory {
					totalCategoryPoints += todo.Points
					if todo.isDone {
						currentCategoryPoints += todo.Points
					}
				}
			} 
			c := fmt.Sprintf("%s %s %d / %d", currentCategory, categorySpaces, currentCategoryPoints, totalCategoryPoints)
			l += categoryStyle.Render(c)
			l += "\n"

		}
		if m.cursor == i {
			cursor = ">"
		}
		if todo.isDone {
		date = todo.CompletionDate.Format("2006-01-02")
		checked = "x"
		s := fmt.Sprintf("%s [%s] %s %s %s    %s", cursor, checked, todo.Text, spacesStr, lipgloss.NewStyle().Foreground(lipgloss.Color("#7F00FF")).Render(strconv.Itoa(todo.Points)), lipgloss.NewStyle().AlignHorizontal(lipgloss.Right).Render(date))
		todoStr = completedTodoStyle.Render(s)
		} else {
			todoStr = fmt.Sprintf("%s [%s] %s %s %s", cursor, checked, todo.Text, spacesStr, lipgloss.NewStyle().Foreground(lipgloss.Color("#7F00FF")).Render(strconv.Itoa(todo.Points)))
		}
		
		if todo.isHidden {
			todoStr = fmt.Sprintf("%s [%s] %s %s %s", cursor, checked, getRandomString(utf8.RuneCountInString(todo.Text)), spacesStr, lipgloss.NewStyle().Foreground(lipgloss.Color("#7F00FF")).Render(strconv.Itoa(todo.Points))) 
		}
		l += todoStr + "\n"
		
	}

	l += "\nPress q to quit."

	m.viewport.SetContent(l)
	return lipgloss.NewStyle().Width(m.width).AlignHorizontal(lipgloss.Center).Render(m.viewport.View())
}
