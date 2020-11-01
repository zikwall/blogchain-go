package user

import (
	"database/sql"
	dbx "github.com/go-ozzo/ozzo-dbx"
	forms2 "github.com/zikwall/blogchain/src/models/user/forms"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func (u UserModel) CreateUser(r *forms2.RegisterForm) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	user := NewUser()

	user.PasswordHash = string(password)
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

	u.AttachProfile(r, user)

	return user, err
}

func (u UserModel) AttachProfile(r *forms2.RegisterForm, user *User) {
	profile := Profile{
		userId:      user.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	status, err := u.Query().Insert("profile", dbx.Params{
		"user_id":      profile.userId,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}).Execute()

	_, err = status.LastInsertId()

	if err != nil {
		panic(err)
	}

	user.Profile = profile
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

func PasswordFirewall(hash string, password string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	return true
}