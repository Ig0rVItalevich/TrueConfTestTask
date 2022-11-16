package service

import (
	"errors"
	refactoring "refactoring/models"
	"refactoring/pkg/repository"
)

var _ User = (*UserService)(nil)

type UserService struct {
	repos repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) Create(input refactoring.User) (int, error) {
	return s.repos.Create(input)
}

func (s *UserService) GetById(id int) (refactoring.User, error) {
	return s.repos.GetById(id)
}

func (s *UserService) GetAll() (refactoring.UserList, error) {
	return s.repos.GetAll()
}

func (s *UserService) Update(id int, input refactoring.UpdateUser) error {
	if !input.Validate() {
		return errors.New("information for update doesn't exist")
	}

	return s.repos.Update(id, input)
}

func (s *UserService) Delete(id int) error {
	return s.repos.Delete(id)
}
