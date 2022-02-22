package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zikwall/blogchain/src/pkg/database"
	"github.com/zikwall/blogchain/src/pkg/exceptions"
	"github.com/zikwall/blogchain/src/services/api/forms"
	"github.com/zikwall/blogchain/src/services/api/utils"

	builder "github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	Repository
}

func UseUserRepository(ctx context.Context, conn *database.Connection) *UserRepository {
	return &UserRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (r *UserRepository) find() *builder.SelectDataset {
	query := r.Connection().Builder()
	return query.Select("user.*").From("user")
}

func (r *UserRepository) CreateUser(form *forms.RegisterForm) (User, error) {
	hash, err := utils.GeneratePasswordHash(form.Password)
	if err != nil {
		return User{}, err
	}

	user := User{
		PasswordHash: string(hash),
		Email:        form.Email,
		Username:     form.Username,
		RegistrationIP: sql.NullString{
			String: form.RegistrationIP,
			Valid:  false,
		},
		CreatedAt: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: false,
		},
	}

	record := builder.Record{
		"password_hash":   user.PasswordHash,
		"email":           user.Email,
		"username":        user.Username,
		"registration_ip": user.RegistrationIP,
		"created_at":      user.ConfirmedAt,
	}

	status, err := r.Connection().
		Builder().
		Insert("user").
		Rows(record).
		Executor().
		ExecContext(r.Context())

	if err != nil {
		return User{}, err
	}
	if user.ID, err = status.LastInsertId(); err != nil {
		return User{}, err
	}
	if err := r.AttachProfile(form, &user); err != nil {
		return User{}, err
	}
	return user, err
}

func (r *UserRepository) AttachProfile(form *forms.RegisterForm, user *User) error {
	profile := Profile{
		userID:      user.ID,
		Name:        form.Name,
		PublicEmail: form.PublicEmail,
		Avatar: sql.NullString{
			String: form.Avatar,
			Valid:  false,
		},
	}

	record := builder.Record{
		"user_id":      profile.userID,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}

	_, err := r.Connection().
		Builder().
		Insert("profile").
		Rows(record).
		Executor().
		ExecContext(r.Context())
	if err != nil {
		return exceptions.ThrowPrivateError(err)
	}
	user.Profile = profile
	return nil
}

func (r *UserRepository) FindByUsernameOrEmail(username, email string) (User, error) {
	user := User{}
	found, err := onUsernameOrMailCondition(r.find(), username, email).
		ScanStructContext(r.context, &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required username or mail was not found"))
	}
	return user, nil
}

func (r *UserRepository) FindByCredentials(credentials string) (User, error) {
	query := withUserProfile(r.find())
	query = onUsernameOrMailCondition(query, credentials, credentials)

	user := User{}
	found, err := query.ScanStructContext(r.Context(), &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required credentials was not found"))
	}
	return user, err
}

func (r *UserRepository) FindByID(id int64) (User, error) {
	user := User{}
	found, err := r.find().
		Where(
			builder.C("id").Eq(id),
		).
		ScanStructContext(r.Context(), &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required ID was not found"))
	}
	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (User, error) {
	query := r.find().Where(
		builder.C("username").Eq(username),
	)
	query = withUserProfile(query)

	user := User{}
	found, err := query.ScanStructContext(r.Context(), &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required username was not found"))
	}
	return user, err
}

func withUserProfile(query *builder.SelectDataset) *builder.SelectDataset {
	query = query.SelectAppend(
		builder.I("profile.name").As(builder.C("profile.name")),
		builder.I("profile.public_email").As(builder.C("profile.public_email")),
		builder.I("profile.avatar").As(builder.C("profile.avatar")),
		builder.I("profile.location").As(builder.C("profile.location")),
		builder.I("profile.description").As(builder.C("profile.description")),
	)
	query = query.LeftJoin(
		builder.T("profile"),
		builder.On(
			builder.I("profile.user_id").Eq(builder.I("user.id")),
		),
	)
	return query
}

func onUsernameOrMailCondition(query *builder.SelectDataset, username, email string) *builder.SelectDataset {
	query = query.Where(
		builder.Or(
			builder.C("username").Eq(username),
			builder.C("email").Eq(email),
		),
	)
	return query
}
