package repository

import (
	refactoring "refactoring/models"
)

type User interface {
	Create(input refactoring.User) (int, error)
	GetById(id int) (refactoring.User, error)
	GetAll() (refactoring.UserList, error)
	Update(id int, input refactoring.UpdateUser) error
	Delete(id int) error
}

type Repository struct {
	User
}

func NewRepository(db string) *Repository {
	return &Repository{
		User: NewUserJson(db),
	}
}
