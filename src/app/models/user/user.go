package user

import (
	"database/sql"
	"github.com/zikwall/blogchain/src/platform/database"
)

type (
	UserModel struct {
		connection *database.Instance
	}
	User struct {
		Id             int64          `db:"id"`
		Username       string         `db:"username"`
		Email          string         `db:"email"`
		PasswordHash   string         `db:"password_hash"`
		ConfirmedAt    sql.NullInt64  `db:"confirmed_at"`
		BlockedAt      sql.NullInt64  `db:"blocked_at"`
		CreatedAt      sql.NullInt64  `db:"created_at"`
		UpdatedAt      sql.NullInt64  `db:"updated_at"`
		RegistrationIp sql.NullString `db:"registration_ip"`

		Profile Profile `db:"profile"`
	}
)

func CreateUserConnection(connection *database.Instance) UserModel {
	return UserModel{
		connection: connection,
	}
}

func (u UserModel) Connection() *database.Instance {
	return u.connection
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
