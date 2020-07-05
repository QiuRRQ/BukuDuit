package models

import "database/sql"

type Shop struct {
	ID          string         `db:"id"`
	FullName    sql.NullString `db:"full_name"`
	BookName    string         `db:"book_name"`
	MobilePhone sql.NullString `db:"mobile_phone"`
	TagLine     sql.NullString `db:"tag_line"`
	Address     sql.NullString `db:"address"`
	Email       sql.NullString `db:"email"`
	Avatar      sql.NullString `db:"avatar"`
	UserID      string         `db:"user_id"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
