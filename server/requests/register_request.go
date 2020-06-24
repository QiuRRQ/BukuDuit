package request

type RegisterRequest struct {
	MobilePhone string `json:"mobile_phone" validate:"required"`
	Pin         string `json:"pin" validate:"required"`
	ShopName    string `json:"shop_name"`
	FullName    string `json:"full_name"`
}
