package viewmodel

type AuthorsVM struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"author_address"`
	City       string `json:"author_city"`
	Province   string `json:"author_province"`
	PostalCode string `json:"author_postal"`
	NoTelp     string `json:"author_no_telp"`
	Created_at string `json:"created_at"`
	UPdated_at string `json:"updated_at`
}
