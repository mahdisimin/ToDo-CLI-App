package service

import (
	"ToDo/entity"
	"errors"
)

type TaskServiceRepository interface {
	PersistTask(entity.Task) (entity.Task, error)
}
type Task struct {
	TaskRepository     TaskServiceRepository
	CategoryRepository CategoryServiceRepository
}

type CreatTaskRequest struct {
	Name                string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreatTaskResponse struct {
	Task     entity.Task
	MetaData string
}

func (t Task) CreateTask(req CreatTaskRequest) (CreatTaskResponse, error) {

	ok := t.CategoryRepository.UserHaveThisCategory(req.AuthenticatedUserID, req.CategoryID)
	if !ok {
		return CreatTaskResponse{}, errors.New("user not have this category")
	}

	task := entity.Task{
		Id:         0,
		Name:       req.Name,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		Status:     "",
		UserID:     req.AuthenticatedUserID,
	}

	persistedTask, _ := t.TaskRepository.PersistTask(task)

	result := CreatTaskResponse{
		Task:     persistedTask,
		MetaData: "",
	}
	return result, nil
}
