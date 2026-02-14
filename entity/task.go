package entity

type Task struct {
	Id         int
	Name       string
	DueDate    string
	CategoryID int
	Status     string
	UserID     int
}

func (task Task) CreatTask() {
}
