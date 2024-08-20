package tasklist

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list     list.Model
	selected int
}

type AddTask string

func openAddTaskForm() tea.Msg {
	return AddTask("Open Add Task Form")
}

func AddNewTask(m *Model, t task.Task) tea.Cmd {
	cmd := m.list.InsertItem(len(m.list.Items()), t)
	return cmd
}

func New() Model {
	newList := list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 40)
	return Model{list: newList}
}

func Update(msg tea.Msg, m *Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if m.selected > 0 {
				m.selected--
				m.list.Select(m.selected)
			}
		case "j", "down":
			if m.selected < len(m.list.Items())-1 {
				m.selected++
				m.list.Select(m.selected)
			}
		case "n":
			return openAddTaskForm
		}
	}
	return nil
}

func View(m *Model) string {
	return m.list.View()
}
