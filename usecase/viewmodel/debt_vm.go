package viewmodel

type DebtVm struct {
	CustomerID     string `json:"customer_id"`
	SubmissionDate string `json:"submission_date"`
	BillDate       string `json:"bill_date"`
	Total          int32  `json:"total"`
	Status         string `json:"status"`
	DebtType       int    `json:"debt_type"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
