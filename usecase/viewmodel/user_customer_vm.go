package viewmodel

type UserCustomerVm struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	MobilePhone string `json:"mobile_phone"`
	Debt        int32  `json:"debt"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
