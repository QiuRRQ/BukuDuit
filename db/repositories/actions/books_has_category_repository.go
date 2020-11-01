package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"time"
)

type BookHasCategory struct {
	DB *sql.DB
}

func NewBookCategoryModel(DB *sql.DB) contracts.IBooksHasCategory {
	return BookHasCategory{DB: DB}
}

func (repository BookHasCategory) Read() (data []models.Books_has_category, err error) {
	statement := `select * from "buku_has_category" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books_has_category{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.BooksId,
			&dataTemp.CategoryID,
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

func (repository BookHasCategory) ReadByID(ID string) (data []models.Books_has_category, err error) {
	statement := `select * from "buku_has_category" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Books_has_category{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.BooksId,
			&dataTemp.CategoryID,
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

func (repository BookHasCategory) Add(body viewmodel.BooksHasCategoryVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "buku_has_category" ("buku_id","category_id","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			body.BookID,
			body.CategoryId,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			body.BookID,
			body.CategoryId,
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasCategory) Edit(body viewmodel.BooksHasCategoryVM) (res string, err error) {
	statement := `update "buku_has_category" set "buku_id"=$1, "category_id"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.BookID,
		body.CategoryId,
		datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasCategory) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "buku_has_category" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasCategory) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "buku_has_category" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BookHasCategory) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "buku_has_category" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
