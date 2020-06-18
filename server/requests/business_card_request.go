package request

type BusinessCardRequest struct {
	FullName    string         `json:"full_name"`
	BookName    string         `json:"book_name"`
	MobilePhone string         `json:"mobile_phone"`
	TagLine     string         `json:"tag_line"`
	Address     string         `json:"address"`
	Email       string         `json:"email"`
	Avatar      string         `json:"avatar"`
	UserID      string         `json:"user_id"`
}
