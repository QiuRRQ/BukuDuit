package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"time"
)

type BorrowCard struct {
	DB *sql.DB
}

func NewBorrowCardModel(DB *sql.DB) contracts.IBorrowCard {
	return BorrowCard{DB: DB}
}

func (repository BorrowCard) ReadBorrowedBook() (data []models.RekapPinjaman, err error) {
	statement := `select "members".name, "members".no_member, "books".title, "borrows_card".jumlah,
	"borrows_card".status, "borrows_card".trans_date, "borrows_card".trans_month, "borrows_card".trans_year 
	from "borrows_card" 
	join "members" on "borrows_card".members_id = "members".id
	join "books" on "borrows_card".books_id = "books".id
	
	where "borrows_card"."deleted_at" is null
	and "borrow_done" like '0' and "status"='peminjaman'`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.RekapPinjaman{}

		err = rows.Scan(
			&dataTemp.MemberName,
			&dataTemp.NoMember,
			&dataTemp.BookTitle,
			&dataTemp.JumlahPinjam,
			&dataTemp.Status,
			&dataTemp.Tgl_Pinjam,
			&dataTemp.Bln_Pinjam,
			&dataTemp.Year_Pinjam,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository BorrowCard) ReadKembalianBook() (data []models.RekapPengembalian, err error) {
	statement := `select "members".name, "members".no_member, "books".title, "borrows_card".jumlah,
	"borrows_card".status, "borrows_card".trans_date, "borrows_card".trans_month, "borrows_card".trans_year 
	from "borrows_card" 
	join "members" on "borrows_card".members_id = "members".id
	join "books" on "borrows_card".books_id = "books".id
	
	where "borrows_card"."deleted_at" is null
	and "borrow_done" like '1' and "status"='peminjaman'`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.RekapPengembalian{}

		err = rows.Scan(
			&dataTemp.MemberName,
			&dataTemp.NoMember,
			&dataTemp.BookTitle,
			&dataTemp.JumlahPinjam,
			&dataTemp.Status,
			&dataTemp.Tgl_Pinjam,
			&dataTemp.Bln_Pinjam,
			&dataTemp.Year_Pinjam,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository BorrowCard) Read() (data []models.BorrowCard, err error) {
	statement := `select * from "borrows_card" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BorrowCard{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.MembersId.String,
			&dataTemp.TransDate,
			&dataTemp.TransMonth,
			&dataTemp.TransYear.String,
			&dataTemp.Status,
			&dataTemp.BookId,
			&dataTemp.Jumlah,
			&dataTemp.BorrowDone,
			&dataTemp.Created_at,
			&dataTemp.Updated_at,
			&dataTemp.Deleted_at,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository BorrowCard) ReadByID(ID string) (data []models.BorrowCard, err error) {
	statement := `select * from "borrows_card" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BorrowCard{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.MembersId.String,
			&dataTemp.TransDate,
			&dataTemp.TransMonth,
			&dataTemp.TransYear.String,
			&dataTemp.Status,
			&dataTemp.BookId,
			&dataTemp.Jumlah,
			&dataTemp.BorrowDone,
			&dataTemp.Created_at,
			&dataTemp.Updated_at,
			&dataTemp.Deleted_at,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

//ini untuk mencari Pinjaman yang bukunya belum dikembalikan
func (repository BorrowCard) ReadByBookID(ID, memberID string) (data []models.BorrowCard, err error) {
	statement := `select * from "borrows_card" where "deleted_at" is null and "books_id"=$1 and "members_id"=$2 
	and "borrow_done" like '0' and "status"='peminjaman'`
	rows, err := repository.DB.Query(statement, ID, memberID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BorrowCard{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.MembersId.String,
			&dataTemp.TransDate,
			&dataTemp.TransMonth,
			&dataTemp.TransYear.String,
			&dataTemp.Status,
			&dataTemp.BookId,
			&dataTemp.Jumlah,
			&dataTemp.BorrowDone,
			&dataTemp.Created_at,
			&dataTemp.Updated_at,
			&dataTemp.Deleted_at,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository BorrowCard) Edit(body viewmodel.BorrowCardVM) (res string, err error) {
	statement := `update "borrows_card" set "members_id"=$1, "trans_date"=$2, "trans_month"=$3, 
	"trans_year"=$4, "status"=$5, "books_id"=$6,"jumlah"=$7, "borrow_done"=$8, "updated_at"=$9 where "id"=$10 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.MembersId,
		body.TransDate,
		body.TransMonth,
		body.TransYear,
		body.Status,
		body.BookId,
		body.Jumlah,
		body.BorrowDone,
		datetime.StrParseToTime(body.Updated_at, time.RFC3339),
		body.Id).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BorrowCard) EditBorrowDone(body viewmodel.BorrowCardVM) (res string, err error) {
	statement := `update "borrows_card" set "borrow_done"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.BorrowDone,
		datetime.StrParseToTime(body.Updated_at, time.RFC3339),
		body.Id).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BorrowCard) Add(body viewmodel.BorrowCardVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "borrows_card" 
	("members_id","trans_date","trans_month","trans_year","status","books_id","jumlah","borrow_done","created_at","updated_at") 
	values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.MembersId),
			str.EmptyString(body.TransDate),
			str.EmptyString(body.TransMonth),
			str.EmptyString(body.TransYear),
			str.EmptyString(body.Status),
			str.EmptyString(body.BookId),
			body.Jumlah,
			body.BorrowDone,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.Updated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.MembersId),
			str.EmptyString(body.TransDate),
			str.EmptyString(body.TransMonth),
			str.EmptyString(body.TransYear),
			str.EmptyString(body.Status),
			str.EmptyString(body.BookId),
			body.Jumlah,
			body.BorrowDone,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.Updated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BorrowCard) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "borrows_card" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BorrowCard) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "borrows_card" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BorrowCard) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "borrows_card" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
