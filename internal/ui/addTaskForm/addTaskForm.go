package addtaskform

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textAreaDescription textarea.Model
	textInputTitle      textinput.Model
}

func New() Model {
	title := textinput.New()
	title.Placeholder = "Enter a title for the new task."
	title.Focus()

	description := textarea.New()
	description.Placeholder = "Enter a description for the new task."

	form := Model{textAreaDescription: description, textInputTitle: title}

	return form
}

func (m Model) CreateTask() tea.Msg {
	task := task.NewTask(task.Pending, m.textInputTitle.Value(), m.textAreaDescription.Value())
	return task
}

/* Reset all inputs to default. */
func Clear(m *Model) {
	m.textInputTitle.Reset()
	m.textAreaDescription.Reset()
	m.textInputTitle.Focus()
}

func Update(msg tea.Msg, m *Model) tea.Cmd {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.textInputTitle.Focused() {
				m.textInputTitle.Blur()
				m.textAreaDescription.Focus()
				return textarea.Blink
			} else {
				return m.CreateTask
			}
		}
	}
	if m.textInputTitle.Focused() {
		m.textInputTitle, cmd = m.textInputTitle.Update(msg)
		return cmd
	} else {
		m.textAreaDescription, cmd = m.textAreaDescription.Update(msg)
		return cmd
	}
}

func View(m Model) string {
	var s string
	s += m.textInputTitle.View()
	s += "\n\n"
	s += m.textAreaDescription.View()

	return s
}
