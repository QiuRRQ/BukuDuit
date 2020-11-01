package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"time"
)

type BookHasBarcode struct {
	DB *sql.DB
}

func NewBarcodeModel(DB *sql.DB) contracts.IBarcode {
	return BookHasBarcode{DB: DB}
}

func (repository BookHasBarcode) Read() (data []models.Books_has_barcode, err error) {
	statement := `select * from "books_barcode" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books_has_barcode{}

		err = rows.Scan(
			&dataTemp.Barcode,
			&dataTemp.BooksId,
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

func (repository BookHasBarcode) ReadByID(Barcode string) (data []models.Books_has_barcode, err error) {
	statement := `select * from "books_barcode" where "deleted_at" is null and "barcode"=$1`
	rows, err := repository.DB.Query(statement, Barcode)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books_has_barcode{}

		err = rows.Scan(
			&dataTemp.Barcode,
			&dataTemp.BooksId,
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

func (repository BookHasBarcode) Add(body viewmodel.BarcodeVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "books_barcode" ("books_id","created_at","updated_at") values($1,$2,$3) returning "barcode"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			body.BooksID,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			body.BooksID,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasBarcode) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "books_barcode" set "updated_at"=$1, "deleted_at"=$2 where "barcode"=$3 returning  "barcode"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasBarcode) CountByPk(ID string) (res int, err error) {
	statement := `select count("barcode") from "books_barcode" where "barcode"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasBarcode) CountBy(column, value string) (res int, err error) {
	statement := `select count("barcode") from "books_barcode" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
