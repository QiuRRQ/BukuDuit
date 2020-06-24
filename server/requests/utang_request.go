package request

type UtangRequest struct {
	CustomerID    		string 	`json:"customer_id"`
	DebtType    		string 	`json:"debt_type"` //pay = customer bayar hutang, debt = customer minta hutang 
	UserCustomerDebt 	int 	`json: "customer_debt"`
	Amount 				int 	`json : "from_id_amount"`

}
