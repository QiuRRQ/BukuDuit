package request

type MemberRequest struct {
	NoMember   string `json:"no_member"`
	Name       string `json:"name"`
	NoTelp     string `json:"no_telp"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Province   string `json:"province"`
	MemberImg  string `json:"member_img"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	BirthMonth string `json:"birth_month"`
	BirthYear  string `json:"birth_year"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}
