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

type Books struct {
	DB *sql.DB
}

func NewBooksModel(DB *sql.DB) contracts.IBooks {
	return Books{DB: DB}
}

func (repository Books) Read() (data []models.Books, err error) {
	statement := `select * from "books" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Title,
			&dataTemp.Publisher_id,
			&dataTemp.Authors_id,
			&dataTemp.Book_img,
			&dataTemp.Stock,
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

func (repository Books) ReadByID(ID string) (data []models.Books, err error) {
	statement := `select * from "books" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Title,
			&dataTemp.Publisher_id,
			&dataTemp.Authors_id,
			&dataTemp.Book_img,
			&dataTemp.Stock,
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

func (repository Books) Edit(body viewmodel.BooksVM) (res string, err error) {
	statement := `update "books" set "title"=$1, "publisher_id"=$2, "authors_id"=$3, "book_img"=$4, "stock"=$5, "updated_at"=$6 where "id"=$7 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.Tittle,
		body.Publisher_id,
		body.Authors_id,
		body.Book_img,
		body.Stock,
		datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Books) EditStock(body viewmodel.BooksVM, tx *sql.Tx) (res string, err error) {
	statement := `update "books" set "stock"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			body.Stock,
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
			body.ID,
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			body.Stock,
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
			body.ID).Scan(&res)
	}

	if err != nil {
		return res, err
	}

	return res, err
}

//end here
func (repository Books) Add(body viewmodel.BooksVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "books" ("title","publisher_id","authors_id","book_img","stock","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.Tittle),
			body.Publisher_id,
			str.EmptyString(body.Authors_id),
			str.EmptyString(body.Book_img),
			body.Stock,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.Tittle),
			body.Publisher_id,
			str.EmptyString(body.Authors_id),
			str.EmptyString(body.Book_img),
			body.Stock,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Books) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "books" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Books) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "books" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Books) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "books" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
