package viewmodel

type BooksDebtVm struct {
	ID             string `json:"id"`
	CustomerID     string `json:"customer_id"`
	SubmissionDate string `json:"submission_date"`
	BillDate       string `json:"bill_date"`
	DebtTotal      int    `json:"debt_total"`
	CreditTotal    int    `json:"credit_total"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
