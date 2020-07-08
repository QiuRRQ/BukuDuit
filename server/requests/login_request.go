package request

type LoginRequest struct {
	MobilePhone string `json:"mobile_phone"`
	PIN         string `json:"pin"`
}
