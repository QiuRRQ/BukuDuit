package request

type UserRequest struct {
	ID          string `json:"id"`
	MobilePhone string `json:"mobile_phone"`
	OldPin      string `json:"old_pin"`
	NewPin      string `json:"new_pin"`
	Name        string `json:"full_name"`
}
