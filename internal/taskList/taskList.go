package tasklist

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	Description string
	Done        bool
}

type Model struct {
	List     list.Model
	Tasks    []Task
	Quitting bool
}

type item string

type ItemDelegate struct{}

func (i item) FilterValue() string { return "" }

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)
	fmt.Fprint(w, str)
}

func (t Task) ToListItem() item {
	doneStr := " "

	if t.Done {
		doneStr = "x"
	}
	s := fmt.Sprintf("[%s] %s", doneStr, t.Description)

	return item(s)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() string {
	if m.Quitting {
		return "Quitting...\n"
	}

	return m.List.View()
}
