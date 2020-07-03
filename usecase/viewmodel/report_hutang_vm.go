package viewmodel

type ReportHutangVm struct {
	ShopID      string       `json:"shop_id"`
	TotalCredit int          `json:"total_credit"`
	TotalDebit  int          `json:"total_debit"`
	ListData    []DebtReport `json:"list_data"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"update_at"`
	DeletedAt   string       `json:"deleted_at"`
}

type DebtReport struct {
	TransactionDate string       `json:"transaction_date"`
	Details         []DebtDetail `json:"details"`
}

type DebtDetail struct {
	ID          string `json:"id"`
	ReferenceID string `json:"reference_id"` //tak perlu tampilin ini.
	Name        string `json:"full_name"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
	Type        string `json:"type"`
}
