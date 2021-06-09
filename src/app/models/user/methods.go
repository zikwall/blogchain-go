package user

import (
	"database/sql"
	"errors"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/app/exceptions"
	"github.com/zikwall/blogchain/src/app/models/user/forms"
	"github.com/zikwall/blogchain/src/app/utils"
	"time"
)

func (m Model) find() *builder.SelectDataset {
	query := m.Connection().Builder()
	return query.Select("user.*").From("user")
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

func (m Model) CreateUser(r *forms.RegisterForm) (User, error) {
	hash, err := utils.GenerateBlogchainPasswordHash(r.Password)

	if err != nil {
		return User{}, err
	}

	user := User{}
	user.PasswordHash = string(hash)
	user.Email = r.Email
	user.Username = r.Username
	user.RegistrationIp = sql.NullString{
		String: r.RegistrationIp,
		Valid:  false,
	}
	user.CreatedAt.Int64 = time.Now().Unix()

	record := builder.Record{
		"password_hash":   user.PasswordHash,
		"email":           user.Email,
		"username":        user.Username,
		"registration_ip": user.RegistrationIp,
		"created_at":      user.ConfirmedAt,
	}

	status, err := m.Connection().
		Builder().
		Insert("user").
		Rows(record).
		Executor().
		ExecContext(m.Context())

	if err != nil {
		return User{}, err
	}

	if user.Id, err = status.LastInsertId(); err != nil {
		return User{}, err
	}

	if err = m.AttachProfile(r, &user); err != nil {
		return User{}, err
	}

	return user, err
}

func (m Model) AttachProfile(r *forms.RegisterForm, user *User) error {
	profile := Profile{
		userId:      user.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	record := builder.Record{
		"user_id":      profile.userId,
		"name":         profile.Name,
		"public_email": profile.PublicEmail,
		"avatar":       profile.Avatar,
	}

	_, err := m.Connection().
		Builder().
		Insert("profile").
		Rows(record).
		Executor().
		ExecContext(m.Context())

	if err != nil {
		return exceptions.NewErrDatabaseAccess(err)
	}

	user.Profile = profile

	return nil
}

func (m Model) FindByUsernameOrEmail(username string, email string) (User, error) {
	user := User{}
	found, err := onUsernameOrMailCondition(m.find(), username, email).
		ScanStructContext(m.context, &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required username or mail was not found"))
	}

	return user, nil
}

func (m Model) FindByCredentials(credentials string) (User, error) {
	query := withUserProfile(m.find())
	query = onUsernameOrMailCondition(query, credentials, credentials)

	user := User{}
	found, err := query.ScanStructContext(m.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required credentials was not found"))
	}

	return user, err
}

func (m Model) FindById(id int64) (User, error) {
	user := User{}
	found, err := m.find().
		Where(
			builder.C("id").Eq(id),
		).
		ScanStructContext(m.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required ID was not found"))
	}

	return user, nil
}

func (m Model) FindByUsername(username string) (User, error) {
	query := m.find().Where(
		builder.C("username").Eq(username),
	)
	query = withUserProfile(query)

	user := User{}
	found, err := query.ScanStructContext(m.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required username was not found"))
	}

	return user, err
}
