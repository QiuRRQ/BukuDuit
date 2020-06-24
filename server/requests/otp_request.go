package request

type OtpRequest struct {
	MobilePhone string `json:"mobile_phone"`
}

type InviteUserRequest struct {
	MobilePhone string `json:"mobile_phone"`
	Type        string `json:"type"`
}

type SendfeedBackRequest struct {
	Message string `json:"message"`
}

type OtpSubmitRequest struct {
	MobilePhone string `json:"mobile_phone" validate:"required"`
	Otp         string `json:"otp" validate:"required"`
	Action      string `json:"action"`
}
