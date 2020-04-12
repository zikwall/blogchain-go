package forms

type RegisterForm struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	RegistrationIp string
}

func (r *RegisterForm) ComparePasswords() bool {
	return r.Password == r.PasswordRepeat
}

func (r *RegisterForm) Validate() bool {
	return r.required() && r.minLengths()
}

func (r *RegisterForm) required() bool {
	return r.Email != "" && r.Username != "" && r.Password != ""
}

func (r *RegisterForm) minLengths() bool {
	return len(r.Username) > 5 && len(r.Password) > 8
}
