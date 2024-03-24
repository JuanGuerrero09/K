package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	View() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
}

type shortAnswerField struct {
	textinput textinput.Model
}

type longAnswerField struct {
	textarea textarea.Model
}



func NewShortAnswerField() *shortAnswerField {
	ti := textinput.New()
	ti.Placeholder = "Your answer here..."
	ti.Focus()
	return &shortAnswerField{ti}
}

func (sa *shortAnswerField) Value() string {
	return sa.textinput.Value()
}

func (sa *shortAnswerField) View() string {
	return sa.textinput.View()
}

func (sa *shortAnswerField) Blur() tea.Msg {
	return sa.textinput.Blur
}

func (sa *shortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textinput, cmd = sa.textinput.Update(msg)
	return sa, cmd
}

func NewLongAnswerField() *longAnswerField {
	ta := textarea.New()
	ta.Placeholder = "Your answer here..."
	ta.Focus()
	return &longAnswerField{ta}
}

func (la *longAnswerField) Value() string {
	return la.textarea.Value()
}

func (la *longAnswerField) View() string {
	return la.textarea.View()
}

func (la *longAnswerField) Blur() tea.Msg {
	return la.textarea.Blur
}

func (la *longAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textarea, cmd = la.textarea.Update(msg)
	return la, cmd
}
