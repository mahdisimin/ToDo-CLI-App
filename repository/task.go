package repository

import (
	"ToDo/entity"
	"fmt"
)

type UserStorage interface {
	Create(entity any) (any, error)
	Read() []entity.User
	Update()
	Delete()
}
type task struct {
	FilePath string
	storage  UserStorage
}

func (t task) PersistTask(task entity.Task) (entity.Task, error) {
	userList := t.storage.Read()
	userCount := len(userList)
	task.Id = userCount + 1

	result, err := t.storage.Create(task)
	if err != nil {
		fmt.Errorf("Error while persisting task: %v", err)
	}
	persistedTask, ok := result.(entity.Task)
	if !ok {
		fmt.Errorf("Error while persisting task: %v", err)
	}
	return persistedTask, nil
}
