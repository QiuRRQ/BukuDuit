package request

type TransactionRequest struct {
	ReferenceID     string `json:"reference_id" validate:"required"`
	TransactionType string `json:"transaction_type" validate:"required"` //pay = customer bayar hutang, debt = customer minta hutang
	ShopID          string `json:"shop_id" validate:"required"`
	Amount          int32  `json:"amount" validate:"required,numeric"`
	TransactionDate string `json:"transaction_date"`
}
