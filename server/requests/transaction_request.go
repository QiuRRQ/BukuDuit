package request

type TransactionRequest struct {
	ID               string `json:"id"`
	Customer_Id      string `json:"customer_id"`
	Amount           string `json:"amount"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	Type             string `json:"type"`
	Transaction_Date string `json:"transaction_date"`
	Created_at       string `json:"created_at"`
	Update_at        string `json:"update_at"`
	Deleted_at       string `json:"deleted_at"`
}
