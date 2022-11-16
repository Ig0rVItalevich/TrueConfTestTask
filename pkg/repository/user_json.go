package repository

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	refactoring "refactoring/models"
	"strconv"
	"time"
)

var _ User = (*UserJson)(nil)

type UserJson struct {
	db string
}

func NewUserJson(db string) *UserJson {
	return &UserJson{db: db}
}

func (r *UserJson) Create(input refactoring.User) (int, error) {
	var store refactoring.UserStore
	data, err := ioutil.ReadFile(r.db)
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(data, &store); err != nil {
		return 0, err
	}

	*store.Increment += 1
	input.CreatedAt = time.Now()

	id := strconv.Itoa(*store.Increment)
	store.List[id] = input

	dataNew, err := json.MarshalIndent(&store, "", "    ")
	if err != nil {
		return 0, err
	}

	if err := ioutil.WriteFile(r.db, dataNew, fs.ModePerm); err != nil {
		return 0, err
	}

	return *store.Increment, nil
}

func (r *UserJson) GetById(id int) (refactoring.User, error) {
	var user refactoring.User
	var store refactoring.UserStore
	data, err := ioutil.ReadFile(r.db)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(data, &store); err != nil {
		return user, err
	}

	idStr := strconv.Itoa(id)
	user, ok := store.List[idStr]
	if !ok {
		return user, errors.New("user does not exist")
	}

	return user, nil
}

func (r *UserJson) GetAll() (refactoring.UserList, error) {
	var users refactoring.UserList
	var store refactoring.UserStore
	data, err := ioutil.ReadFile(r.db)
	if err != nil {
		return users, err
	}
	if err := json.Unmarshal(data, &store); err != nil {
		return users, err
	}

	return store.List, nil
}

func (r *UserJson) Update(id int, input refactoring.UpdateUser) error {
	var store refactoring.UserStore
	data, err := ioutil.ReadFile(r.db)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &store); err != nil {
		return err
	}

	idStr := strconv.Itoa(id)
	user, ok := store.List[idStr]
	if !ok {
		return errors.New("user does not exist")
	}

	if input.DisplayName != nil {
		user.DisplayName = *input.DisplayName
	}
	if input.Email != nil {
		user.Email = *input.Email
	}

	store.List[idStr] = user

	dataNew, err := json.MarshalIndent(&store, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(r.db, dataNew, fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func (r *UserJson) Delete(id int) error {
	var store refactoring.UserStore
	data, err := ioutil.ReadFile(r.db)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &store); err != nil {
		return err
	}

	idStr := strconv.Itoa(id)
	_, ok := store.List[idStr]
	if !ok {
		return errors.New("user does not exist")
	}

	delete(store.List, idStr)

	dataNew, err := json.MarshalIndent(&store, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(r.db, dataNew, fs.ModePerm); err != nil {
		return err
	}

	return nil
}
