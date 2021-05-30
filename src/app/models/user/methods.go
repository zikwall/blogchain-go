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

func (self Model) Find() *builder.SelectDataset {
	return self.Connection().Builder().Select("user.*").From("user")
}

func (self Model) WithProfile(query *builder.SelectDataset) *builder.SelectDataset {
	return query.
		SelectAppend(
			builder.I("profile.name").As(builder.C("profile.name")),
			builder.I("profile.public_email").As(builder.C("profile.public_email")),
			builder.I("profile.avatar").As(builder.C("profile.avatar")),
			builder.I("profile.location").As(builder.C("profile.location")),
			builder.I("profile.description").As(builder.C("profile.description")),
		).
		LeftJoin(
			builder.T("profile"),
			builder.On(
				builder.I("profile.user_id").Eq(builder.I("user.id")),
			),
		)
}

func (self Model) onCredentialsCondition(query *builder.SelectDataset, username, email string) *builder.SelectDataset {
	return query.
		Where(
			builder.Or(
				builder.C("username").Eq(username),
				builder.C("email").Eq(email),
			),
		)
}

func (self Model) CreateUser(r *forms.RegisterForm) (User, error) {
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

	insert := self.Connection().
		Builder().
		Insert("user").
		Rows(
			builder.Record{
				"password_hash":   user.PasswordHash,
				"email":           user.Email,
				"username":        user.Username,
				"registration_ip": user.RegistrationIp,
				"created_at":      user.ConfirmedAt,
			},
		).Executor()

	status, err := insert.ExecContext(self.Context())

	if err != nil {
		return User{}, err
	}

	if user.Id, err = status.LastInsertId(); err != nil {
		return User{}, err
	}

	if err = self.AttachProfile(r, &user); err != nil {
		return User{}, err
	}

	return user, err
}

func (self Model) AttachProfile(r *forms.RegisterForm, user *User) error {
	profile := Profile{
		userId:      user.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	insert := self.Connection().
		Builder().Insert("profile").
		Rows(
			builder.Record{
				"user_id":      profile.userId,
				"name":         profile.Name,
				"public_email": profile.PublicEmail,
				"avatar":       profile.Avatar,
			},
		).Executor()

	if _, err := insert.ExecContext(self.Context()); err != nil {
		return exceptions.NewErrDatabaseAccess(err)
	}

	user.Profile = profile

	return nil
}

func (self Model) FindByUsernameOrEmail(username string, email string) (User, error) {
	user := User{}
	found, err := self.onCredentialsCondition(self.Find(), username, email).ScanStructContext(self.context, &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required username or mail was not found"))
	}

	return user, nil
}

func (self Model) FindByCredentials(credentials string) (User, error) {
	user := User{}
	query := self.WithProfile(self.Find())
	query = self.onCredentialsCondition(query, credentials, credentials)

	found, err := query.ScanStructContext(self.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required credentials was not found"))
	}

	return user, err
}

func (self Model) FindById(id int64) (User, error) {
	user := User{}
	found, err := self.Find().Where(builder.C("id").Eq(id)).ScanStructContext(self.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required ID was not found"))
	}

	return user, nil
}

func (self Model) FindByUsername(username string) (User, error) {
	user := User{}
	query := self.Find().Where(builder.C("username").Eq(username))
	query = self.WithProfile(query)

	found, err := query.ScanStructContext(self.Context(), &user)

	if err != nil {
		return user, exceptions.NewErrDatabaseAccess(err)
	} else if !found {
		return user, exceptions.NewErrApplicationLogic(errors.New("user with the required username was not found"))
	}

	return user, err
}
