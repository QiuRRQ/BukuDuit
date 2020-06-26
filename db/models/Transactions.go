package models

import (
	"database/sql"
)

type Transactions struct {
	ID              string         `db:"id"`
	ReferenceID     string         `db:"reference_id"`
	Name            string         `db:"full_name"`
	IDShop          string         `db:"shop_id"`
	Amount          sql.NullInt32  `db:"amount"`
	Description     sql.NullString `db:"description"`
	Image           sql.NullString `db:"image"`
	Type            string         `db:"type"`
	TransactionDate sql.NullString `db:"transaction_date"`
	CreatedAt       string         `db:"created_at"`
	UpdatedAt       sql.NullString `db:"update_at"`
	DeletedAt       sql.NullString `db:"deleted_at"`
	Status          bool           `db:"status"`
}
