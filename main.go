package main

import (
	"fmt"
	"os"

	tasklist "tasks/internal/taskList"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := tasklist.NewModel(80, 40)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
