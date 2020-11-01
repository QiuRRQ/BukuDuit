package request

type PublisherRequest struct {
	Id             string `json:"author_id"`
	AuthorName     string `json:"author_name"`
	AuthorAddress  string `json:"author_address"`
	AuthorCity     string `json:"author_city"`
	AuthorProvince string `json:"author_province"`
	PostalCode     string `json:"author_postal_code"`
	NoTelp         string `json:"author_no_telp"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
	Deleted_at     string `json:"deleted_at"`
}
