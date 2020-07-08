package models

import "database/sql"

type BooksDebt struct {
	ID             string         `db:"id"`
	CustomerID     string         `db:"customer_id"`
	SubmissionDate string         `db:"submission_date"`
	BillDate       sql.NullString `db:"bill_date"`
	DebtTotal      int            `db:"debt_total"`
	CreditTotal    int            `db:"credit_total"`
	Status         sql.NullString `db:"status"`
	CreatedAt      string         `db:"created_at"`
	UpdatedAt      sql.NullString `db:"updated_at"`
	DeletedAt      sql.NullString `db:"deleted_at"`
}
