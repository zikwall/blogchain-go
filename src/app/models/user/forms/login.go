package forms

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginForm) Validate() bool {
	return l.Username != "" && l.Password != ""
}
