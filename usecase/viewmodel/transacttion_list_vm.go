package viewmodel

type TransactionListVm struct {
	ShopID      string     `json:"shop_id"`
	TotalCredit int        `json:"total_credit"`
	TotalDebit  int        `json:"total_debit"`
	ListData    []DataList `json:"list_data"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"update_at"`
	DeletedAt   string     `json:"deleted_at"`
}

type WeeklyListVM struct {
	Start			string			`json:"start"`
	End				string			`json:"end"`
	Debit			int				`json:"debit"`
	Credit			int				`json:"credit"`
	Details			[]DataDetails 	`json:"details"`
}

type DataList struct {
	TransactionDate  string        `json:"transaction_date"`
	DateAmountCredit int           `json:"date_credit_amount"`
	DateAmountDebet  int           `json:"date_debet_amount"`
	Details          []DataDetails `json:"details"`
}

type DataDetails struct {
	ID              string `json:"id"`
	TransactionDate string `json:"transaction_date"`
	ReferenceID     string `json:"reference_id"` //tak perlu tampilin ini.
	Name            string `json:"full_name"`
	Description     string `json:"description"`
	Amount          int32  `json:"amount"`
	Type            string `json:"type"`
}
