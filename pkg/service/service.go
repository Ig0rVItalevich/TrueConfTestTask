package service

import (
	refactoring "refactoring/models"
	"refactoring/pkg/repository"
)

type User interface {
	Create(input refactoring.User) (int, error)
	GetById(id int) (refactoring.User, error)
	GetAll() (refactoring.UserList, error)
	Update(id int, input refactoring.UpdateUser) error
	Delete(id int) error
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
	}
}
