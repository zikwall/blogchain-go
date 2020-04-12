package user

import (
	"database/sql"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/zikwall/blogchain/di"
	"github.com/zikwall/blogchain/models/user/forms"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func CreateUser(r *forms.RegisterForm) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	u := &User{
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
		CreatedAt: 0,
		UpdatedAt: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		RegistrationIp: sql.NullString{
			String: r.RegistrationIp,
			Valid:  false,
		},
	}

	u.PasswordHash = string(password)
	u.Email = r.Email
	u.Username = r.Username
	u.RegistrationIp = sql.NullString{
		String: r.RegistrationIp,
		Valid:  false,
	}
	u.CreatedAt = time.Now().Unix()

	status, err := di.DI().Database.Query().Insert("user", dbx.Params{
		"password_hash":   u.PasswordHash,
		"email":           u.Email,
		"username":        u.Username,
		"registration_ip": u.RegistrationIp,
		"created_at":      u.ConfirmedAt,
	}).Execute()

	u.Id, err = status.LastInsertId()

	return u, err
}

func FindByUsernameOrEmail(username string, email string) (*User, error) {
	user := &User{
		Id:           0,
		Username:     "",
		Email:        "",
		PasswordHash: "",
	}

	err := di.DI().Database.Query().
		Select("*").
		From("user").
		Where(dbx.Or(dbx.HashExp{"username": username}, dbx.HashExp{"email": email})).
		One(&user)

	return user, err
}

func FindByCredentials(credentials string) (*User, error) {
	user := &User{
		Id:           0,
		Username:     "",
		Email:        "",
		PasswordHash: "",
	}

	err := di.DI().Database.Query().
		Select("*").
		From("user").
		Where(dbx.Or(dbx.HashExp{"username": credentials}, dbx.HashExp{"email": credentials})).
		One(&user)

	return user, err
}

func FindById(id int64) (*User, error) {
	user := &User{
		Id:           0,
		Username:     "",
		Email:        "",
		PasswordHash: "",
	}

	err := di.DI().Database.Query().
		Select("*").
		From("user").
		Where(dbx.HashExp{"id": id}).
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
