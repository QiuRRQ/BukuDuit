package request

type UtangRequest struct {
	Reference_id string `json:"reference_id"`
	DebtType     string `json:"debt_type"` //pay = customer bayar hutang, debt = customer minta hutang
	Shop_id      int    `json: "shop_id"`
	Amount       int    `json : "from_id_amount"`
}
