package request

type BorrowCardRequest struct {
	BorrowDone string `json:"borrow_done"`
	MembersId  string `json:"members_id"`
	TransDate  string `json:"trans_date"`
	TransMonth string `json:"trans_month"`
	TransYear  string `json:"trans_year"`
	Status     string `json:"status"`
	BookId     string `json:"books_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
	Id         string `json:"id"`
	Jumlah     int64  `json:"jumlah"`
}
