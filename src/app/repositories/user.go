package repositories

import "database/sql"

type User struct {
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

type Profile struct {
	userId      int64
	Name        string         `db:"name"`
	PublicEmail string         `db:"public_email"`
	Avatar      sql.NullString `db:"avatar"`
	Location    sql.NullString `db:"location"`
	Status      sql.NullString `db:"status"`
	Description sql.NullString `db:"description"`
}

type PublicProfile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
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
