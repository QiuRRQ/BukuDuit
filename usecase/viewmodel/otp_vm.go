package viewmodel

type OtpVm struct {
	MobilePhone string `json:"phone"`
	ExpiredDate string `json:"expired_date"`
	Otp         string `json:"otp"`
}

// InvalidOtpCounterVM ....
type InvalidOtpCounterVM struct {
	Counter int `json:"counter"`
}
