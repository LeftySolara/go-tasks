package tasklist

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/* Task Model
*
* A Bubble Tea model used to display the task list.
 */

type Model struct {
	err      error
	list     list.Model
	tasks    []task.Task
	quitting bool
	loaded   bool
	focused  int
}

func NewModel(width, height int) *Model {
	newList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	newList.SetItems([]list.Item{
		task.NewTask(task.Pending, "Example Task", "This is an example task."),
		task.NewTask(task.InProgress, "Example Task 2", "This is an example task."),
		task.NewTask(task.Done, "Example Task 3", "This is an example task."),
	})

	return &Model{
		list: newList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	// Don't render to stdout if a "quit" key has been pressed.
	// This prevents artifacts from appearing on the command line after program exit.
	if m.quitting {
		return ""
	}

	if m.loaded {
		listView := m.list.View()
		return listView
	} else {
		return "loading..."
	}
}
