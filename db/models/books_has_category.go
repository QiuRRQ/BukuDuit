package models

import "database/sql"

type Books_has_category struct {
	Id         string         `db:"id"`
	BooksId    string         `db:"buku_id"`
	CategoryID string         `db:"category_id"`
	Created_at string         `db:"created_at"`
	Updated_at string         `db:"updated_at"`
	Deleted_at sql.NullString `db:"deleted_at"`
}
