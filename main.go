package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type task struct {
	description string
	done        bool
}

type model struct {
	list     list.Model
	tasks    []task
	quitting bool
}

type listItem string

type listItemDelegate struct{}

func (d listItemDelegate) Height() int                             { return 1 }
func (d listItemDelegate) Spacing() int                            { return 0 }
func (d listItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d listItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(listItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fmt.Fprint(w, str)
}

func (i listItem) FilterValue() string { return "" }

func (t task) toListItem() listItem {
	doneStr := " "

	if t.done {
		doneStr = "x"
	}
	s := fmt.Sprintf("[%s] %s", doneStr, t.description)

	return listItem(s)
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

	return m.list.View()
}

func main() {
	tasks := []task{
		{description: "Example", done: false},
		{description: "Example 2", done: true},
	}

	items := []list.Item{
		tasks[0].toListItem(),
		tasks[1].toListItem(),
	}

	l := list.New(items, listItemDelegate{}, 20, 14)
	m := model{quitting: false, tasks: tasks, list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
