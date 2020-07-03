package viewmodel

type UserCustomerDebetCreditVm struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Type        string `json:"type"`
	Amount      int32  `json:"amount"`
	MobilePhone string `json:"mobile_phone"`
}
