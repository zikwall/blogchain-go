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
	CreatedAt      sql.NullInt64
	UpdatedAt      sql.NullInt64
	RegistrationIp sql.NullString

	Profile Profile
}

func NewUser() *User {
	return &User{
		Id:           0,
		Username:     "",
		Email:        "",
		PasswordHash: "",
		ConfirmedAt: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		BlockedAt: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		CreatedAt: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		UpdatedAt: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		RegistrationIp: sql.NullString{
			String: "",
			Valid:  false,
		},
		Profile: NewProfile(),
	}
}

type PublicUser struct {
	Id       int64         `json:"id"`
	Username string        `json:"username"`
	Profile  PublicProfile `json:"profile"`
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
		Profile: PublicProfile{
			Name:        u.Profile.Name,
			Email:       u.Profile.PublicEmail,
			Avatar:      u.Profile.Avatar.String,
			Location:    u.Profile.Location.String,
			Status:      u.Profile.Status.String,
			Description: u.Profile.Description.String,
		},
	}
}
