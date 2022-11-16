package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	refactoring "refactoring/models"
)

func NewJsonStore(filePath string) (string, error) {
	var store refactoring.UserStore
	
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(data, &store); err != nil {
		return "", err
	}
	if store.Increment == nil || store.List == nil {
		return "", errors.New("invalid store file")
	}

	return filePath, err
}
