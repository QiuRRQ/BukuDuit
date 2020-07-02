package viewmodel

type PaymentAccountVm struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ShopID        string `json:"shop_id"`
	PaymentNumber string `json:"payment_number"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

type ListPaymentAcc struct {
	ListAccPayment []PaymentAccountVm `json:"list_account"`
}
