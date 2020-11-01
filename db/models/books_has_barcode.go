package models

import "database/sql"

type Books_has_barcode struct {
	Barcode    string         `db:"barcode"`
	BooksId    string         `db:"books_id"`
	Created_at string         `db:"created_at"`
	Updated_at string         `db:"updated_at"`
	Deleted_at sql.NullString `db:"deleted_at"`
}
