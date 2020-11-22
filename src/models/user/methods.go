package user

import (
	"database/sql"
	dbx "github.com/go-ozzo/ozzo-dbx"
	forms2 "github.com/zikwall/blogchain/src/models/user/forms"
	"github.com/zikwall/blogchain/src/utils"
	"time"
)

type UserProfile struct {
	User    `db:"user"`
	Profile `db:"profile"`
}

func (u UserModel) CreateUser(r *forms2.RegisterForm) (*User, error) {
	hash, err := utils.GenerateBlogchainPasswordHash(r.Password)

	if err != nil {
		return nil, err
	}

	user := NewUser()

	user.PasswordHash = string(hash)
	user.Email = r.Email
	user.Username = r.Username
	user.RegistrationIp = sql.NullString{
		String: r.RegistrationIp,
		Valid:  false,
	}
	user.CreatedAt.Int64 = time.Now().Unix()

	status, err := u.Query().Insert("user", dbx.Params{
		"password_hash":   user.PasswordHash,
		"email":           user.Email,
		"username":        user.Username,
		"registration_ip": user.RegistrationIp,
		"created_at":      user.ConfirmedAt,
	}).Execute()

	user.Id, err = status.LastInsertId()

	err = u.AttachProfile(r, user)

	return user, err
}

func (u UserModel) AttachProfile(r *forms2.RegisterForm, user *User) error {
	profile := Profile{
		userId:      user.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	_, err := u.Query().Insert("profile", dbx.Params{
		"user_id":      profile.userId,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}).Execute()

	if err != nil {
		return err
	}

	user.Profile = profile

	return nil
}

func (u UserModel) FindByUsernameOrEmail(username string, email string) (*User, error) {
	user := NewUser()

	err := u.Query().
		Select("*").
		From("user").
		Where(dbx.Or(dbx.HashExp{"username": username}, dbx.HashExp{"email": email})).
		One(&user)

	return user, err
}

func (u UserModel) FindByCredentials(credentials string) (*User, error) {
	user := NewUser()

	err := u.Query().
		Select("user.*", "p.name as profile.name", "p.public_email as profile.public_email", "p.avatar as profile.avatar").
		From("user").
		LeftJoin("profile p", dbx.NewExp("p.user_id=user.id")).
		Where(dbx.Or(dbx.HashExp{"user.username": credentials}, dbx.HashExp{"user.email": credentials})).
		One(&user)

	return user, err
}

func (u UserModel) FindById(id int64) (*User, error) {
	user := NewUser()

	err := u.Query().
		Select("user.*", "p.name as profile.name", "p.public_email as profile.public_email", "p.avatar as profile.avatar").
		From("user").
		LeftJoin("profile p", dbx.NewExp("p.user_id=user.id")).
		Where(dbx.HashExp{"user.id": id}).
		One(&user)

	return user, err
}

func (u UserModel) FindByUsername(username string) (*User, error) {
	user := NewUser()

	err := u.Query().
		Select("user.username", "user.id",
			"p.name as profile.name",
			"p.public_email as profile.public_email",
			"p.avatar as profile.avatar",
			"p.location as profile.location",
			"p.status as profile.status",
			"p.description as profile.description",
		).
		From("user").
		LeftJoin("profile p", dbx.NewExp("p.user_id=user.id")).
		Where(dbx.HashExp{"username": username}).
		One(&user)

	return user, err
}
