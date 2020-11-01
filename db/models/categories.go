package models

import "database/sql"

type Categories struct {
	Id         string         `db:"id"`
	Name       string         `db:"name"`
	Created_at string         `db:"created_at"`
	Updated_at string         `db:"updated_at"`
	Deleted_at sql.NullString `db:"deleted_at"`
}
