package models

import "database/sql"

type Publishers struct {
	Id         string         `db:"id"`
	Name       sql.NullString `db:"name"`
	Address    string         `db:"address"`
	City       string         `db:"city"`
	Province   sql.NullString `db:"province"`
	PostalCode sql.NullString `db:"postal_code"`
	NoTelp     sql.NullString `db:"no_telp"`
	Created_at string         `db:"created_at"`
	Updated_at string         `db:"updated_at"`
	Deleted_at sql.NullString `db:"deleted_at"`
}
