package forms

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

func (r *RegisterForm) ComparePasswords() bool {
	return r.Password == r.PasswordRepeat
}

func (r *RegisterForm) Validate() bool {
	return r.required() && r.minLengths()
}

func (r *RegisterForm) required() bool {
	return r.Email != "" && r.Username != "" && r.Password != "" && r.Name != ""
}

func (r *RegisterForm) minLengths() bool {
	return len(r.Username) > 5 && len(r.Password) > 8
}
