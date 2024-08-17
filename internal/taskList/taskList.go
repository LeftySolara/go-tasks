package tasklist

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Task List
*
* A Bubble Tea model used to display the task list.
 */

type TaskList struct {
	err      error
	list     list.Model
	quitting bool
	loaded   bool
	selected int
}

func NewTaskList(width, height int) *TaskList {
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

	return &TaskList{
		list:     newList,
		selected: 0,
	}
}

func (taskList TaskList) SelectPrev() {
	if taskList.selected > 0 {
		taskList.selected--
		taskList.list.Select(taskList.selected)
	}
}

func (taskList TaskList) SelectNext() {
	if taskList.selected < len(taskList.list.Items())-1 {
		taskList.selected++
		taskList.list.Select(taskList.selected)
	}
}

/* list.Item interface */

func (taskList TaskList) Init() tea.Cmd {
	return nil
}

func (taskList TaskList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		taskList.list.SetWidth(msg.Width)
		taskList.list.SetHeight(msg.Height)
		if !taskList.loaded {
			taskList.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			taskList.quitting = true
			return taskList, tea.Quit
		case "up", "k":
			taskList.SelectPrev()
		case "down", "j":
			taskList.SelectNext()
		}
	}
	var cmd tea.Cmd
	taskList.list, cmd = taskList.list.Update(msg)
	return taskList, cmd
}

func (taskList TaskList) View() string {
	// Don't render to stdout if a "quit" key has been pressed.
	// This prevents artifacts from appearing on the command line after program exit.
	if taskList.quitting {
		return ""
	}

	if taskList.loaded {
		listView := taskList.list.View()
		return listView
	} else {
		return "loading..."
	}
}
