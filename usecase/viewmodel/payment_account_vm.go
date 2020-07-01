package viewmodel

type PaymentAccountVm struct {
	ID            string `json:"id"`
	ShopID        string `json:"shop_id"`
	Name          string `json:"name"`
	PaymentNumber string `json:"payment_number"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}
