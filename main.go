package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialModel := NewModel()
	if _, err := tea.NewProgram(initialModel, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type model struct {
	TaskListSelectedIndex int
	TextInputTitle        textinput.Model
	TextAreaDescription   textarea.Model
	TaskList              list.Model
	Quitting              bool
	AddingTask            bool
}

// TODO: Add new task to the task list.
func (m model) CreateTask() tea.Msg {
	return nil
}

func NewModel() tea.Model {
	m := model{}
	m.TaskList = list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 40)
	m.TextInputTitle = textinput.New()
	m.TextAreaDescription = textarea.New()
	m.Quitting = false
	m.AddingTask = false

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit.
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if m.AddingTask {
			if k == "esc" || k == "ctrl+c" {
				m.Quitting = true
				return m, tea.Quit
			}
		} else {
			if k == "q" || k == "esc" || k == "ctrl+c" {
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.AddingTask {
		return updateAddTaskForm(msg, m)
	}
	return updateTaskList(msg, m)
}

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n Exiting...\n\n"
	}
	if m.AddingTask {
		s = addTaskFormView(m)
	} else {
		s = taskListView(m)
	}

	return s
}

func updateTaskList(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	if m.AddingTask {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if m.TaskListSelectedIndex < len(m.TaskList.Items()) {
				m.TaskListSelectedIndex++
				m.TaskList.Select(m.TaskListSelectedIndex)
			}
		case "j", "down":
			if m.TaskListSelectedIndex > 0 {
				m.TaskListSelectedIndex--
				m.TaskList.Select(m.TaskListSelectedIndex)
			}
		case "n":
			m.TextInputTitle.Reset()
			m.TextAreaDescription.Reset()
			m.TextInputTitle.Focus()
			m.AddingTask = true
		}
	}
	return m, nil
}

func updateAddTaskForm(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.TextInputTitle.Focused() {
				m.TextInputTitle.Blur()
				m.TextAreaDescription.Focus()
				return m, textarea.Blink
			} else {
				m.AddingTask = false
				return m, m.CreateTask
			}
		}
	}
	if m.TextInputTitle.Focused() {
		m.TextInputTitle, cmd = m.TextInputTitle.Update(msg)
		return m, cmd
	} else {
		m.TextAreaDescription, cmd = m.TextAreaDescription.Update(msg)
		return m, cmd
	}
}

func taskListView(m model) string {
	return m.TaskList.View()
}

func addTaskFormView(m model) string {
	var s string
	s += m.TextInputTitle.View()
	s += "\n"
	s += m.TextAreaDescription.View()

	return s
}
