package addtaskform

import (
	"tasks/internal/task"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AddTaskForm struct {
	textAreaDescription textarea.Model
	textInputTitle      textinput.Model
}

func NewAddTaskForm() *AddTaskForm {
	form := &AddTaskForm{}
	form.textInputTitle = textinput.New()
	form.textAreaDescription = textarea.New()
	form.textInputTitle.Focus()

	return form
}

func (form AddTaskForm) CreateTask() tea.Msg {
	task := task.NewTask(
		task.Pending,
		form.textInputTitle.Value(),
		form.textAreaDescription.Value())

	return task
}

/* MODEL INTERFACE IMPLEMENTATION */
func (form AddTaskForm) Init() tea.Cmd {
	return nil
}

func (form AddTaskForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return form, tea.Quit
		case "enter":
			if form.textInputTitle.Focused() {
				form.textInputTitle.Blur()
				form.textAreaDescription.Focus()
				return form, textarea.Blink
			} else {
				return form, form.CreateTask
			}
		}
	}

	if form.textInputTitle.Focused() {
		form.textInputTitle, cmd = form.textInputTitle.Update(msg)
		return form, cmd
	} else {
		form.textAreaDescription, cmd = form.textAreaDescription.Update(msg)
		return form, cmd
	}
}

func (form AddTaskForm) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, form.textInputTitle.View(), form.textAreaDescription.View())
}
