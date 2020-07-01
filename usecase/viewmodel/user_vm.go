package viewmodel

type UserVm struct {
	ID             string           `json:"id"`
	MobilePhone    string           `json:"mobile_phone"`
	Name           string           `json:"name"`
	PaymentDetails []ListPaymentAcc `json:"payment_details"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
	DeletedAt      string           `json:"deleted_at"`
}
