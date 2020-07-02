package request

type PaymentAccountRequest struct {
	ID    				string         `json:"id"`
	ShopID    			string         `json:"shop_id"`
	Name 				string         `json:"name"`
	PaymentNumber     	string         `json:"payment_number"`
}