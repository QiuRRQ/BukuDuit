package viewmodel

type RekapPinjamanVM struct {
	MemberName   string `json:"member_name"`
	NoMember     string `json:"no_member"`
	BookTitle    string `json:"book_title"`
	JumlahPinjam string `json:"jumlah_dipinjam"`
	Status       string `json:"status_peminjaman"`
	Tgl_Pinjam   string `json:"tgl_pinjam"`
	Bln_Pinjam   string `json:"bln_pinjam"`
	Year_Pinjam  string `json:"year_pinjam"`
}
