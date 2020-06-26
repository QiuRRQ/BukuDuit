package viewmodel

type BusinessCardVm struct {
	ID                  string `json:"id"`
	FullName            string `json:"full_name"`
	BookName            string `json:"book_name"`
	MobilePhone         string `json:"mobile_phone"`
	TagLine             string `json:"tag_line"`
	Address             string `json:"address"`
	Email               string `json:"email"`
	Avatar              string `json:"avatar"`
	TotalCustomerCredit int32  `json:"total_credit"`
	TotalOwnerCredit    int32  `json:"total_owner_credit"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	DeletedAt           string `json:"deleted_at"`
}
