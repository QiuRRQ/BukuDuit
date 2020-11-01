package request

type BookRequest struct {
	Title        string `json:"title"`
	Publisher_id string `json:"publisher_id"`
	Author_id    string `json:"author_id"`
	Book_img     string `json:"book_img"`
	Stock        int64  `json:"stock"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
	Deleted_at   string `json:"deleted_at"`
}
