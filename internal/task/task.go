package task

/* Task
*
* Data structure representing a single task.
* Implements the list.Item and list.DefaultItem interfaces from Bubble Tea.
 */

type status int

/* TASK STATUSES */
const (
	Pending status = iota
	InProgress
	Done
)

type Task struct {
	title       string
	description string
	status      status
}

func NewTask(status status, title, description string) Task {
	return Task{title: title, description: description, status: status}
}

/* list.Item interface */
func (t Task) FilterValue() string {
	return t.title
}

/* list.DefaultItem interface */
func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}
