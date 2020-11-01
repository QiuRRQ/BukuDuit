package viewmodel

type BooksVM struct {
	ID           string `json:"id"`
	Tittle       string `json:"title"`
	Publisher_id string `json:"publisher_id"`
	Authors_id   string `json:"authors_id"`
	Book_img     string `json:"book_img"`
	Stock        int64  `json:"stock"`
	Created_at   string `json:"created_at"`
	UPdated_at   string `json:"updated_at`
}
