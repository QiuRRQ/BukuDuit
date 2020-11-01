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

type Publishers struct {
	DB *sql.DB
}

func NewPublishersModel(DB *sql.DB) contracts.IPublishers {
	return Publishers{DB: DB}
}

func (repository Publishers) Read() (data []models.Publishers, err error) {
	statement := `select * from "publishers" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Publishers{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Name,
			&dataTemp.Address,
			&dataTemp.City,
			&dataTemp.Province,
			&dataTemp.PostalCode,
			&dataTemp.NoTelp,
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

func (repository Publishers) ReadByID(ID string) (data []models.Publishers, err error) {
	statement := `select * from "publishers" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Publishers{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.Name,
			&dataTemp.Address,
			&dataTemp.City,
			&dataTemp.Province,
			&dataTemp.PostalCode,
			&dataTemp.NoTelp,
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

func (repository Publishers) Edit(body viewmodel.PublishersVM) (res string, err error) {
	statement := `update "publishers" set "name"=$1, "address"=$2, "city"=$3, "province"=$4, "postal_code"=$5, "no_telp"=$6, "updated_at"=$7 where "id"=$8 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.Name,
		body.Address,
		body.City,
		body.Province,
		body.PostalCode,
		body.NoTelp,
		datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Publishers) Add(body viewmodel.PublishersVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "publishers" ("name","address","city","province","postal_code","no_telp","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7,$8) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.Name),
			str.EmptyString(body.Address),
			str.EmptyString(body.City),
			str.EmptyString(body.Province),
			str.EmptyString(body.PostalCode),
			str.EmptyString(body.NoTelp),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.Name),
			str.EmptyString(body.Address),
			str.EmptyString(body.City),
			str.EmptyString(body.Province),
			str.EmptyString(body.PostalCode),
			str.EmptyString(body.NoTelp),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.UPdated_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Publishers) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "publishers" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Publishers) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "publishers" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Publishers) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "publishers" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
