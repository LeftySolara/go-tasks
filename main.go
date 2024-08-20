package main

import (
	"fmt"
	"os"
	"tasks/internal/task"
	addtaskform "tasks/internal/ui/addTaskForm"
	tasklist "tasks/internal/ui/taskList"

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
	TaskList    tasklist.Model
	AddTaskForm addtaskform.Model
	Quitting    bool
	AddingTask  bool
}

func NewModel() tea.Model {
	m := model{}
	m.TaskList = tasklist.New()
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
		cmd := tasklist.AddNewTask(&m.TaskList, newTask)

		return m, cmd
	}

	if _, ok := msg.(tasklist.AddTask); ok {
		addtaskform.Clear(&m.AddTaskForm)
		m.AddingTask = true
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.AddingTask {
		cmd := addtaskform.Update(msg, &m.AddTaskForm)
		return m, cmd
	}
	return m, tasklist.Update(msg, &m.TaskList)
}

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n Exiting...\n\n"
	}
	if m.AddingTask {
		s = addtaskform.View(m.AddTaskForm)
	} else {
		s = tasklist.View(&m.TaskList)
	}

	return s
}
