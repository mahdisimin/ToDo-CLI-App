package service

import (
	"ToDo/entity"
	"errors"
)

type TaskServiceRepository interface {
	UserHaveThisCategory(userId int, categoryId int) bool
	PersistTask(entity.Task) (entity.Task, error)
}
type task struct {
	Repository TaskServiceRepository
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

func (t task) CreateTask(req CreatTaskRequest) (CreatTaskResponse, error) {

	ok := t.Repository.UserHaveThisCategory(req.AuthenticatedUserID, req.CategoryID)
	if !ok {
		return CreatTaskResponse{}, errors.New("User Not Have This Cateogry")
	}
	task := entity.Task{
		Id:         0,
		Name:       req.Name,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		Status:     "",
		UserID:     req.AuthenticatedUserID,
	}

	persistedTask, _ := t.Repository.PersistTask(task)

	result := CreatTaskResponse{
		Task:     persistedTask,
		MetaData: "",
	}
	return result, nil
}
