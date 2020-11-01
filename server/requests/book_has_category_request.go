package request

type BookHasCategoryRequest struct {
	Id         string `json:"id"`
	BookId     string `json:"book_id"`
	CategoryId string `json:"category_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}
