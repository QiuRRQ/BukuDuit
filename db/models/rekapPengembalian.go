package models

type RekapPengembalian struct {
	MemberName   string `db:"name"`
	NoMember     string `db:"no_member"`
	BookTitle    string `db:"title"`
	JumlahPinjam string `db:"jumlah"`
	Status       string `db:"status"`
	Tgl_Pinjam   string `db:"trans_date"`
	Bln_Pinjam   string `db:"trans_month"`
	Year_Pinjam  string `db:"trans_year"`
}
