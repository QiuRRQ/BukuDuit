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

type Members struct {
	DB *sql.DB
}

func NewMembersModel(DB *sql.DB) contracts.IMembers {
	return Members{DB: DB}
}

func (repository Members) Read() (data []models.Members, err error) {
	statement := `select * from "members" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Members{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.NoMember,
			&dataTemp.Name,
			&dataTemp.NoTelp,
			&dataTemp.Address,
			&dataTemp.City,
			&dataTemp.Province,
			&dataTemp.Member_img,
			&dataTemp.Gender,
			&dataTemp.BirthDate,
			&dataTemp.BirthMonth,
			&dataTemp.BirthYear,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository Members) ReadByID(ID string) (data []models.Members, err error) {
	statement := `select * from "members" where "deleted_at" is null and "id"=$1`
	rows, err := repository.DB.Query(statement, ID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Members{}

		err = rows.Scan(
			&dataTemp.Id,
			&dataTemp.NoMember,
			&dataTemp.Name,
			&dataTemp.NoTelp,
			&dataTemp.Address,
			&dataTemp.City,
			&dataTemp.Province,
			dataTemp.Member_img,
			&dataTemp.Gender,
			&dataTemp.BirthDate,
			&dataTemp.BirthMonth,
			&dataTemp.BirthYear,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository Members) Edit(body viewmodel.MembersVM) (res string, err error) {
	statement := `update "members" set "no_member"=$1, "name"=$2, "no_telp"=$3, "address"=$4, "city"=$5, 
	"province"=$6, "member_img"=$7, "gender"=$8, "birth_date"=$9,
	"birth_month"=$10, "birth_year"=$11, "updated_at"=$12 where "id"=$13 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.NoMember,
		body.Name,
		body.NoTelp,
		body.Address,
		body.City,
		body.Province,
		body.Member_IMG,
		body.Gender,
		body.BirthDate,
		body.BirthMonth,
		body.BirthYear,
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Members) Add(body viewmodel.MembersVM, tx *sql.Tx) (res string, err error) {
	statement := `insert into "members" ("no_member","name","no_telp","address","city","province",
	"member_img","gender","birth_date","birth_month","birth_year","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.NoMember),
			str.EmptyString(body.Name),
			str.EmptyString(body.NoTelp),
			str.EmptyString(body.Address),
			str.EmptyString(body.City),
			str.EmptyString(body.Province),
			str.EmptyString(body.Member_IMG),
			str.EmptyString(body.Gender),
			str.EmptyString(body.BirthDate),
			str.EmptyString(body.BirthMonth),
			str.EmptyString(body.BirthYear),
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.NoMember),
			str.EmptyString(body.Name),
			str.EmptyString(body.NoTelp),
			str.EmptyString(body.Address),
			str.EmptyString(body.City),
			str.EmptyString(body.Province),
			str.EmptyString(body.Member_IMG),
			str.EmptyString(body.Gender),
			str.EmptyString(body.BirthDate),
			str.EmptyString(body.BirthMonth),
			str.EmptyString(body.BirthYear),
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Members) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "members" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Members) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "members" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository Members) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "books" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
