package entity

type Task struct {
	Id           int
	TaskName     string
	TaskDueDate  string
	TaskCategory string
	TaskStatus   string
	UserID       int
}

func (task Task) CreatTask() {
}
