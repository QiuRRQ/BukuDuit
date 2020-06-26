package viewmodel

type DetailsHutangVm struct {
	ID          string     `json:"id"`
	ReferenceID string     `json:"reference_id"`
	Name        string     `json:"full_name"`
	ShopID      string     `json:"shop_id"`
	TotalCredit int        `json:"total_credit"`
	TotalDebit  int        `json:"total_debit"`
	ListData    []DebtList `json:"list_data"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"update_at"`
	DeletedAt   string     `json:"deleted_at"`
}

type DebtList struct {
	TransactionDate string   `json:"transaction_date"`
	Details         []Detail `json:"details"`
}

type Detail struct {
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
	Type        string `json:"type"`
}
