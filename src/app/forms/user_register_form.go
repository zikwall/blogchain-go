package forms

import "errors"

type RegisterForm struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	RegistrationIp string
	Name           string `json:"name"`
	PublicEmail    string `json:"public_email"`
	Avatar         string `json:"avatar"`
}

func (r *RegisterForm) ComparePasswords() error {
	if r.Password == r.PasswordRepeat {
		return errors.New("passwords don't match")
	}

	return nil
}

func (r *RegisterForm) Validate() error {
	if err := r.required(); err != nil {
		return err
	}

	if err := r.minLengths(); err != nil {
		return err
	}

	return r.ComparePasswords()
}

func (r *RegisterForm) required() error {
	if r.Email != "" && r.Username != "" && r.Password != "" && r.Name != "" {
		return nil
	}

	return errors.New("email, username or password is empty")
}

func (r *RegisterForm) minLengths() error {
	if len(r.Username) > 5 && len(r.Password) > 8 {
		return nil
	}

	return errors.New("username or password must be at least in length 5 and 8 characters")
}
