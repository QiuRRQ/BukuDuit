package viewmodel

type TransactionVm struct {
	ID              string `json:"id"`
	ReferenceID     string `json:"reference_id"`
	Name            string `json:"full_name"`
	ShopID          string `json:"shop_id"`
	Amount          int32  `json:"amount"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	Type            string `json:"type"`
	TransactionDate string `json:"transaction_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"update_at"`
	DeletedAt       string `json:"deleted_at"`
}
