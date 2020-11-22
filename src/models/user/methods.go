package user

import (
	"database/sql"
	builder "github.com/doug-martin/goqu/v9"
	"github.com/zikwall/blogchain/src/models/user/forms"
	"github.com/zikwall/blogchain/src/utils"
	"time"
)

func (self UserModel) Find() *builder.SelectDataset {
	return self.QueryBuilder().Select("user.*").From("user")
}

func (self UserModel) WithProfile(query *builder.SelectDataset) *builder.SelectDataset {
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

func (self UserModel) onCredentialsCondition(query *builder.SelectDataset, username, email string) *builder.SelectDataset {
	return query.
		Where(
			builder.Or(
				builder.C("username").Eq(username),
				builder.C("email").Eq(email),
			),
		)
}

func (self UserModel) CreateUser(r *forms.RegisterForm) (User, error) {
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

	insert := self.QueryBuilder().Insert("user").Rows(
		builder.Record{
			"password_hash":   user.PasswordHash,
			"email":           user.Email,
			"username":        user.Username,
			"registration_ip": user.RegistrationIp,
			"created_at":      user.ConfirmedAt,
		},
	).Executor()

	status, err := insert.Exec()
	user.Id, err = status.LastInsertId()

	err = self.AttachProfile(r, &user)

	return user, err
}

func (self UserModel) AttachProfile(r *forms.RegisterForm, user *User) error {
	profile := Profile{
		userId:      user.Id,
		Name:        r.Name,
		PublicEmail: r.PublicEmail,
		Avatar: sql.NullString{
			String: r.Avatar,
			Valid:  false,
		},
	}

	insert := self.QueryBuilder().Insert("profile").Rows(
		builder.Record{
			"user_id":      profile.userId,
			"name":         profile.Name,
			"public_email": profile.PublicEmail,
			"avatar":       profile.Avatar,
		},
	).Executor()

	if _, err := insert.Exec(); err != nil {
		return err
	}

	user.Profile = profile

	return nil
}

func (self UserModel) FindByUsernameOrEmail(username string, email string) (User, error) {
	user := User{}
	query := self.Find()
	query = self.onCredentialsCondition(query, username, email)

	_, err := query.ScanStruct(&user)

	return user, err
}

func (self UserModel) FindByCredentials(credentials string) (User, error) {
	user := User{}
	query := self.Find()
	query = self.WithProfile(query)
	query = self.onCredentialsCondition(query, credentials, credentials)

	_, err := query.ScanStruct(&user)

	return user, err
}

func (self UserModel) FindById(id int64) (User, error) {
	user := User{}
	query := self.Find().Where(builder.C("id").Eq(id))

	_, err := query.ScanStruct(&user)

	return user, err
}

func (self UserModel) FindByUsername(username string) (User, error) {
	user := User{}
	query := self.Find().Where(builder.C("username").Eq(username))
	query = self.WithProfile(query)

	_, err := query.ScanStruct(&user)

	return user, err
}
