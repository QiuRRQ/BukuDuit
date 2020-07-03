package models

import "database/sql"

type PaymentAccount struct {
	ID            string         `db:"id"`
	ShopID        string         `db:"shop_id"`
	Name          sql.NullString `db:"account_name"`
	OwnerName     sql.NullString `db:"owner_name"`
	PaymentNumber string         `db:"payment_number"`
	CreatedAt     string         `db:"created_at"`
	UpdatedAt     sql.NullString `db:"updated_at"`
	DeletedAt     sql.NullString `db:"deleted_at"`
}
