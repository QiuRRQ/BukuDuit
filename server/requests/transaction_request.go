package request

type TransactionRequest struct {
	ID              string `json:"id"`
	CustomerID      string `json:"customer_id"`
	Amount          int32  `json:"amount"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	Type            string `json:"type"`
	TransactionDate string `json:"transaction_date"`
}
