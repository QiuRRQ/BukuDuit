package request

type BarcodeRequest struct {
	BooksID    string `json:"book_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}
