package viewmodel

type UserCustomerDebetCreditVm struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	MobilePhone string `json:"phone_number"`
	Type        string `json:"type"`
	Amount      int32  `json:"amount"`
}
