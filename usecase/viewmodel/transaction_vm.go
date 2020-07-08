package viewmodel

type TransactionVm struct {
	ID              string `json:"id"`
	ReferenceID     string `json:"reference_id"`
	Name            string `json:"full_name"`
	ShopID          string `json:"shop_id"`
	CustomerID      string `json:"customer_id"`
	Amount          int32  `json:"amount"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	Type            string `json:"type"`
	BooksDebtID     string `json:"books_debt_id"`
	BooksTransID    string `json:"books_transaction_id"`
	TransactionDate string `json:"transaction_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"update_at"`
	DeletedAt       string `json:"deleted_at"`
}
