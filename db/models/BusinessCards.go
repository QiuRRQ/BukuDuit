package models

import "database/sql"

type BusinessCards struct {
	ID          string         `db:"id"`
	FullName    string         `db:"full_name"`
	BookName    string         `db:"book_name"`
	MobilePhone string         `db:"mobile_phone"`
	TagLine     string         `db:"tag_line"`
	Address     string         `db:"address"`
	Email       string         `db:"email"`
	Avatar      string         `db:"avatar"`
	UserID      string         `db:"user_id"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
