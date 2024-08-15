package tasklist

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Task Model
*
* A Bubble Tea model used to display the task list.
 */

type Model struct {
	err      error
	list     list.Model
	quitting bool
	loaded   bool
	selected int
}

func NewModel(width, height int) *Model {
	delegate := list.NewDefaultDelegate()

	// Set colors for selected task
	focusedColor := lipgloss.Color("13")
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(focusedColor).
		BorderLeftForeground(focusedColor)

	newList := list.New([]list.Item{}, delegate, width, height)
	newList.SetItems([]list.Item{
		task.NewTask(task.Pending, "Example Task", "This is an example task."),
		task.NewTask(task.InProgress, "Example Task 2", "This is an example task."),
		task.NewTask(task.Done, "Example Task 3", "This is an example task."),
	})

	return &Model{
		list:     newList,
		selected: 0,
	}
}

func (m Model) SelectPrev() {
	if m.selected > 0 {
		m.selected--
		m.list.Select(m.selected)
	}
}

func (m Model) SelectNext() {
	if m.selected < len(m.list.Items())-1 {
		m.selected++
		m.list.Select(m.selected)
	}
}

/* list.Item interface */

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
		case "up", "k":
			m.SelectPrev()
		case "down", "j":
			m.SelectNext()
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
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
