package models

import "database/sql"

type BorrowCard struct {
	Id         string         `db:"id"`
	MembersId  sql.NullString `db:"members_id"`
	TransDate  string         `db:"trans_date"`
	TransMonth string         `db:"trans_month"`
	TransYear  sql.NullString `db:"trans_year"`
	Status     sql.NullString `db:"status"`
	BookId     sql.NullString `db:"books_id"`
	Jumlah     int64          `db:"jumlah"`
	BorrowDone string         `db:"borrows_card"`
	Created_at string         `db:"created_at"`
	Updated_at string         `db:"updated_at"`
	Deleted_at sql.NullString `db:"deleted_at"`
}
