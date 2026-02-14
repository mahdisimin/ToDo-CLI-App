package repository

import (
	"ToDo/contracts"
	"ToDo/entity"
)

type UserRepository interface {
	FindByName(useName, password string) (entity.User, error)
}

type UserService struct {
	repo    UserRepository
	storage contracts.Storage
}

func (service *UserService) Login(userName, password string) (entity.User, error) {
	user, err := service.repo.FindByName(userName, password)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
