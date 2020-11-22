package user

import "database/sql"

type Profile struct {
	userId      int64          `db:"user_id"`
	Name        string         `db:"name"`
	PublicEmail string         `db:"public_email"`
	Avatar      sql.NullString `db:"avatar"`
	Location    sql.NullString `db:"location"`
	Status      sql.NullString `db:"status"`
	Description sql.NullString `db:"description"`
}

type PublicProfile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func NewProfile() Profile {
	return Profile{
		userId:      0,
		Name:        "",
		PublicEmail: "",
		Avatar: sql.NullString{
			String: "",
			Valid:  false,
		},
		Location: sql.NullString{
			String: "",
			Valid:  false,
		},
		Status: sql.NullString{
			String: "",
			Valid:  false,
		},
		Description: sql.NullString{
			String: "",
			Valid:  false,
		},
	}
}
