package user

import (
	"database/sql"
	builder "github.com/doug-martin/goqu/v9"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/zikwall/blogchain/src/models"
	forms2 "github.com/zikwall/blogchain/src/models/user/forms"
	"github.com/zikwall/blogchain/src/utils"
	"time"
)

func (self UserModel) Find() *builder.SelectDataset {
	return models.QueryBuilder().Select("user.*").From("user")
}

func (self UserModel) WithProfile(query *builder.SelectDataset) *builder.SelectDataset {
	return query.
		SelectAppend(
			builder.I("profile.name").As(builder.C("profile.name")),
			builder.I("profile.public_email").As(builder.C("profile.public_email")),
			builder.I("profile.avatar").As(builder.C("profile.avatar")),
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

func (u UserModel) CreateUser(r *forms2.RegisterForm) (User, error) {
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

	status, err := u.Query().Insert("user", dbx.Params{
		"password_hash":   user.PasswordHash,
		"email":           user.Email,
		"username":        user.Username,
		"registration_ip": user.RegistrationIp,
		"created_at":      user.ConfirmedAt,
	}).Execute()

	user.Id, err = status.LastInsertId()

	err = u.AttachProfile(r, &user)

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
