package main

import (
	"fmt"
	"os"

	tasklist "tasks/internal/taskList"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tasks := []tasklist.Task{
		{Description: "Example", Done: false},
		{Description: "Example 2", Done: true},
	}

	items := []list.Item{
		tasks[0].ToListItem(),
		tasks[1].ToListItem(),
	}

	l := list.New(items, tasklist.ItemDelegate{}, 20, 14)
	m := tasklist.Model{Quitting: false, Tasks: tasks, List: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
