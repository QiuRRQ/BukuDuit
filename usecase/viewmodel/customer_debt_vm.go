package viewmodel

type CustomerDebtVm struct {

	Debt  CustomerDebtDetailVm `json:"debt"`
	Total int32                `json:"total"`
}

type CustomerDebtDetailVm struct {
	Type   string `json:"type"`
	Amount int32  `json:"amount"`
}
