package viewmodel

type MembersVM struct {
	ID         string `json:"id"`
	NoMember   string `json:"customer_id"`
	Name       string `json:"barang_id"`
	NoTelp     string `json:"penjualan_qty"`
	Address    string `json:"subtotal"`
	City       string `json:"member_city"`
	Province   string `json:"member_province"`
	Member_IMG string `json:"member_img"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	BirthMonth string `json:"birth_month"`
	BirthYear  string `json:"birth_year"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
