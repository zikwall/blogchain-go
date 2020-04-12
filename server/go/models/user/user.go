package user

import (
	"database/sql"
)

type User struct {
	Id             int64
	Username       string
	Email          string
	PasswordHash   string
	ConfirmedAt    sql.NullInt64
	BlockedAt      sql.NullInt64
	CreatedAt      int64
	UpdatedAt      sql.NullInt64
	RegistrationIp sql.NullString

	Profile Profile
}

type Profile struct {
	UserId      int
	Name        string
	PublicEmail string
	Avatar      string
}

type PublicUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *User) GetId() int64 {
	return u.Id
}

func (u *User) Exist() bool {
	return u.Id > 0
}

func (u *User) IsGuest() bool {
	return !u.Exist()
}

func (u *User) Properties() PublicUser {
	return PublicUser{
		Id:       u.Id,
		Username: u.Username,
	}
}
