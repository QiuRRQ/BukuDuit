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

type Category struct {
	DB *sql.DB
}

func NewCategoryModel(DB *sql.DB) contracts.ICategory {
	return Category{DB: DB}
}

func (repository Category) Read() (data []models.Categories, err error) {
	statement := `select * from "category" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Categories{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Name,
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

func (repository Category) ReadByID(ID string) (data []models.Categories, err error) {
	statement := `select * from "category" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Categories{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Name,
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

func (repository Category) Edit(body viewmodel.CategoriesVM) (res string, err error) {
	statement := `update "category" set "name"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.Name,
		datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Category) Add(body viewmodel.CategoriesVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "category" ("name","created_at","updated_at") values($1,$2,$3) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.Name),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.Name),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Category) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "category" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Category) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "category" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Category) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "category" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
