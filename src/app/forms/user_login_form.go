package forms

import "errors"

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginForm) Validate() error {
	if l.Username != "" && l.Password != "" {
		return nil
	}

	return errors.New("username or password is empty")
}
