package request

type PaymentAccountRequest struct {
	ID            string `json:"id"`
	ShopID        string `json:"shop_id"`
	AccountName   string `json:"account_name"`
	OwnerName     string `json:"owner_namer"`
	PaymentNumber string `json:"payment_number"`
}
