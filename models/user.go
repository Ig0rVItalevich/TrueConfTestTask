package refactoring

import (
	"net/http"
	"time"
)

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

func (u *User) Bind(r *http.Request) error { return nil }

type UpdateUser struct {
	CreatedAt   *time.Time `json:"created_at"`
	DisplayName *string    `json:"display_name"`
	Email       *string    `json:"email"`
}

func (u *UpdateUser) Validate() bool {
	if u.DisplayName == nil && u.Email == nil {
		return false
	}

	return true
}

func (u *UpdateUser) Bind(r *http.Request) error { return nil }

type (
	UserList  map[string]User
	UserStore struct {
		Increment *int     `json:"increment"`
		List      UserList `json:"list"`
	}
)
