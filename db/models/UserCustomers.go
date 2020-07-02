package models

import "database/sql"

type UserCustomers struct {
	ID          string         `db:"id"`
	FullName    string         `db:"full_name"`
	MobilePhone string         `db:"mobile_phone"`
	BusinessID  string         `db:"business_id"`
	PaymentDate sql.NullString `db:"payment_date"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
