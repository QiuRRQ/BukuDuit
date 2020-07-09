package models

import (
	"database/sql"
)

type Transactions struct {
	ID              string         `db:"id"`
	ReferenceID     string         `db:"reference_id"`
	Weekly          sql.NullString `db:"weekly"`
	CategoryID      sql.NullString `db:"category_id"`
	IDShop          string         `db:"shop_id"`
	Name            sql.NullString `db:"full_name"`
	BooksDeptID     sql.NullString `db:"books_debt_id"`
	BooksTransID    sql.NullString `db:"books_transaction_id"`
	Amount          sql.NullInt32  `db:"amount"`
	Description     sql.NullString `db:"description"`
	Image           sql.NullString `db:"image"`
	Type            string         `db:"type"`
	TransactionDate sql.NullString `db:"transaction_date"`
	CustomerID      sql.NullString `db:"customer_id"`
	CreatedAt       string         `db:"created_at"`
	UpdatedAt       sql.NullString `db:"update_at"`
	DeletedAt       sql.NullString `db:"deleted_at"`
	Status          bool           `db:"status"`
}
