package request

type BooksTransactionRequest struct {
	ID          string `json:"id"`
	ShopID      string `json:"shop_id"`
	DebtTotal   int    `json:"debt_total"`
	CreditTotal int    `json:"credit_total"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
