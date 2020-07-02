package models

import "database/sql"

type BooksTransaction struct {
	ID          string         `db:"id"`
	ShopID      sql.NullString `db:"shop_id"`
	DebtTotal   int            `db:"debt_total"`
	CreditTotal int            `db:"credit_total"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
