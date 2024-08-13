package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type task struct {
	description string
	done        bool
}

type model struct {
	tasks    []task
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "Quitting...\n"
	}

	output := ""

	for _, task := range m.tasks {
		var doneStr string
		if task.done {
			doneStr = "x"
		} else {
			doneStr = " "
		}

		output += fmt.Sprintf("[%s] %s\n", doneStr, task.description)
	}
	return output
}

func main() {
	tasks := []task{
		{description: "Example", done: false},
		{description: "Example 2", done: true},
	}
	m := model{quitting: false, tasks: tasks}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
