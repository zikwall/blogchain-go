package repositories

import (
	"context"
	"database/sql"
	"errors"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/forms"
	"github.com/zikwall/blogchain/src/app/utils"
	"github.com/zikwall/blogchain/src/platform/database"
	"time"
)

type UserRepository struct {
	Repository
}

func UseUserRepository(ctx context.Context, conn *database.Connection) UserRepository {
	return UserRepository{
		Repository{connection: conn, context: ctx},
	}
}

func (ur UserRepository) find() *builder.SelectDataset {
	query := ur.Connection().Builder()
	return query.Select("user.*").From("user")
}

func (ur UserRepository) CreateUser(r *forms.RegisterForm) (User, error) {
	hash, err := utils.GenerateBlogchainPasswordHash(r.Password)

	if err != nil {
		return User{}, err
	}

	user := User{}
	user.PasswordHash = string(hash)
	user.Email = r.Email
	user.Username = r.Username
	user.RegistrationIP = sql.NullString{
		String: r.RegistrationIP,
		Valid:  false,
	}
	user.CreatedAt.Int64 = time.Now().Unix()

	record := builder.Record{
		"password_hash":   user.PasswordHash,
		"email":           user.Email,
		"username":        user.Username,
		"registration_ip": user.RegistrationIP,
		"created_at":      user.ConfirmedAt,
	}

	status, err := ur.Connection().
		Builder().
		Insert("user").
		Rows(record).
		Executor().
		ExecContext(ur.Context())

	if err != nil {
		return User{}, err
	}

	if user.ID, err = status.LastInsertId(); err != nil {
		return User{}, err
	}

	if err := ur.AttachProfile(r, &user); err != nil {
		return User{}, err
	}

	return user, err
}

func (ur UserRepository) AttachProfile(r *forms.RegisterForm, user *User) error {
	profile := Profile{
		userID:      user.ID,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	record := builder.Record{
		"user_id":      profile.userID,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}

	_, err := ur.Connection().
		Builder().
		Insert("profile").
		Rows(record).
		Executor().
		ExecContext(ur.Context())

	if err != nil {
		return exceptions.ThrowPrivateError(err)
	}

	user.Profile = profile

	return nil
}

func (ur UserRepository) FindByUsernameOrEmail(username, email string) (User, error) {
	user := User{}
	found, err := onUsernameOrMailCondition(ur.find(), username, email).
		ScanStructContext(ur.context, &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required username or mail was not found"))
	}

	return user, nil
}

func (ur UserRepository) FindByCredentials(credentials string) (User, error) {
	query := withUserProfile(ur.find())
	query = onUsernameOrMailCondition(query, credentials, credentials)

	user := User{}
	found, err := query.ScanStructContext(ur.Context(), &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required credentials was not found"))
	}

	return user, err
}

func (ur UserRepository) FindByID(id int64) (User, error) {
	user := User{}
	found, err := ur.find().
		Where(
			builder.C("id").Eq(id),
		).
		ScanStructContext(ur.Context(), &user)

	if err != nil {
		return user, exceptions.ThrowPrivateError(err)
	} else if !found {
		return user, exceptions.ThrowPublicError(errors.New("user with the required ID was not found"))
	}

	return user, nil
}

func (ur UserRepository) FindByUsername(username string) (User, error) {
	query := ur.find().Where(
		builder.C("username").Eq(username),
	)
	query = withUserProfile(query)

	user := User{}
	found, err := query.ScanStructContext(ur.Context(), &user)

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
