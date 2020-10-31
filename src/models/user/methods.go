package user

import (
	"database/sql"
	dbx "github.com/go-ozzo/ozzo-dbx"
	di2 "github.com/zikwall/blogchain/src/di"
	forms2 "github.com/zikwall/blogchain/src/models/user/forms"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func CreateUser(r *forms2.RegisterForm) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	u := NewUser()

	u.PasswordHash = string(password)
	u.Email = r.Email
	u.Username = r.Username
	u.RegistrationIp = sql.NullString{
		String: r.RegistrationIp,
		Valid:  false,
	}
	u.CreatedAt.Int64 = time.Now().Unix()

	status, err := di2.DI().Database.Query().Insert("user", dbx.Params{
		"password_hash":   u.PasswordHash,
		"email":           u.Email,
		"username":        u.Username,
		"registration_ip": u.RegistrationIp,
		"created_at":      u.ConfirmedAt,
	}).Execute()

	u.Id, err = status.LastInsertId()

	AttachProfile(r, u)

	return u, err
}

func AttachProfile(r *forms2.RegisterForm, u *User) {
	profile := Profile{
		userId:      u.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	status, err := di2.DI().Database.Query().Insert("profile", dbx.Params{
		"user_id":      profile.userId,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}).Execute()

	_, err = status.LastInsertId()

	if err != nil {
		panic(err)
	}

	u.Profile = profile
}

func FindByUsernameOrEmail(username string, email string) (*User, error) {
	user := NewUser()

	err := di2.DI().Database.Query().
		Select("*").
		From("user").
		Where(dbx.Or(dbx.HashExp{"username": username}, dbx.HashExp{"email": email})).
		One(&user)

	return user, err
}

func FindByCredentials(credentials string) (*User, error) {
	user := NewUser()

	err := di2.DI().Database.Query().
		Select("user.*", "p.name as profile.name", "p.public_email as profile.public_email", "p.avatar as profile.avatar").
		From("user").
		LeftJoin("profile p", dbx.NewExp("p.user_id=user.id")).
		Where(dbx.Or(dbx.HashExp{"user.username": credentials}, dbx.HashExp{"user.email": credentials})).
		One(&user)

	return user, err
}

func FindById(id int64) (*User, error) {
	user := NewUser()

	err := di2.DI().Database.Query().
		Select("user.*", "p.name as profile.name", "p.public_email as profile.public_email", "p.avatar as profile.avatar").
		From("user").
		LeftJoin("profile p", dbx.NewExp("p.user_id=user.id")).
		Where(dbx.HashExp{"user.id": id}).
		One(&user)

	return user, err
}

func FindByUsername(username string) (*User, error) {
	user := NewUser()

	err := di2.DI().Database.Query().
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
