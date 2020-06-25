package models

import (
	"database/sql"
)

type Transactions struct {
	ID              string         `db:"id"`
	ReferenceID     string         `db:"reference_id"`
	IDShop          string         `db:"shop_id"`
	Amount          sql.NullInt32  `db:"amount"`
	Description     string         `db:"description"`
	Image           string         `db:"image"`
	Type            string         `db:"type"`
	TransactionDate sql.NullString `db:"transaction_date"`
	CreatedAt       string         `db:"created_at"`
	UpdatedAt       sql.NullString `db:"update_at"`
	DeletedAt       sql.NullString `db:"deleted_at"`
}
