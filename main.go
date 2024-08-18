package main

import (
	"fmt"
	"os"
	"tasks/internal/task"
	addtaskform "tasks/internal/ui/addTaskForm"

	"github.com/charmbracelet/bubbles/list"
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
	TaskList              list.Model
	AddTaskForm           addtaskform.Model
	TaskListSelectedIndex int
	Quitting              bool
	AddingTask            bool
}

func NewModel() tea.Model {
	m := model{}
	m.TaskList = list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 40)
	m.AddTaskForm = addtaskform.New()
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

	if msg, ok := msg.(task.Task); ok {
		newTask := msg
		m.AddingTask = false
		return m, m.TaskList.InsertItem(len(m.TaskList.Items()), newTask)
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.AddingTask {
		cmd := addtaskform.Update(msg, &m.AddTaskForm)
		return m, cmd
	}
	return updateTaskList(msg, m)
}

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n Exiting...\n\n"
	}
	if m.AddingTask {
		s = addtaskform.View(m.AddTaskForm)
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
			addtaskform.Clear(&m.AddTaskForm)
			m.AddingTask = true
		}
	}
	return m, nil
}

func taskListView(m model) string {
	return m.TaskList.View()
}
