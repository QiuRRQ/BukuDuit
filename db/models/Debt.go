package models

import "database/sql"

type Debt struct {
	ID             string         `json:"id"`
	CustomerID     string         `json:"customer_id"`
	SubmissionDate string         `json:"submission_date"`
	BillDate       sql.NullString `json:"bill_date"`
	Total          int32          `json:"total"`
	Status         string         `json:"status"`
	Type           int            `json:"type"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      sql.NullString `json:"updated_at"`
	DeletedAt      sql.NullString `json:"deleted_at"`
}
