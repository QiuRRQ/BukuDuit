package request

//user customer request struct
type UserCustomerRequest struct {
	FullName    string `json:"full_name"`
	MobilePhone string `json:"mobile_phone"`
	ShopID      string `json:"shop_id"`
}
