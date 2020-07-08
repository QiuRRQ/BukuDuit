package models

import "database/sql"

type Users struct {
	ID          string         `db:"id"`
	MobilePhone string         `db:"mobile_phone"`
	Pin         string         `db:"pin"`
	Name        sql.NullString `db:"full_name"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
