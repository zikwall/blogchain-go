package user

import "database/sql"

type Profile struct {
	userId      int64
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
