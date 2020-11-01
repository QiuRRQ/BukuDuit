package models

import "database/sql"

type Books struct {
	Id           string         `db:"id"`
	Title        sql.NullString `db:"title"`
	Publisher_id string         `db:"publisher_id"`
	Authors_id   string         `db:"authors_id"`
	Book_img     sql.NullString `db:"book_img"`
	Stock        sql.NullInt64  `db:"stock"`
	Created_at   string         `db:"created_at"`
	Updated_at   string         `db:"updated_at"`
	Deleted_at   sql.NullString `db:"deleted_at"`
}
