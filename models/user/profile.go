package user

import "database/sql"

type Profile struct {
	userId      int64
	Name        string
	PublicEmail string
	Avatar      sql.NullString
	Location    sql.NullString
	Status      sql.NullString
	Description sql.NullString
}

type PublicProfile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
