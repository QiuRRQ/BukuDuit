package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserModel(DB *sql.DB) contracts.IUserRepository {
	return UserRepository{DB: DB}
}

func (repository UserRepository) Browse(search, order, sort string, limit, offset int) (data []models.Users, count int, err error) {
	statement := `select * from "users" where "mobile_phone" like $1 and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, search, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Users{}
		err = rows.Scan(&dataTemp.ID, &dataTemp.MobilePhone, &dataTemp.Pin, &dataTemp.CreatedAt, &dataTemp.UpdatedAt, &dataTemp.DeletedAt)
		if err != nil {
			return data, count, err
		}
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "users" where "mobile_phone" like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, search).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository UserRepository) ReadByPk(ID string) (data models.Users, err error) {
	statement := `select * from "users" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&data.ID, &data.MobilePhone, &data.Pin, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository UserRepository) ReadBy(column, value string) (data models.Users, err error) {
	statement := `select * from "users" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&data.ID, &data.MobilePhone, &data.Pin, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository UserRepository) Edit(input viewmodel.UserVm) (res string, err error) {
	panic("implement me")
}

func (repository UserRepository) EditPin(ID, pin, updatedAt string) (res string, err error) {
	statement := `update "users" set "pin"=$1, "updated_at"=$2 where "id"=$1 returning "id"`
	err = repository.DB.QueryRow(statement, pin, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) Add(input viewmodel.UserVm, pin string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "users" ("mobile_phone","pin","created_at") values($1,$2,$3)`
	_, err = tx.Exec(statement, input.MobilePhone, pin, datetime.StrParseToTime(input.CreatedAt, time.RFC3339))
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "users" set "updated_at"=$1, "deleted_at"=$2 where "id"=$1 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339)).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "users" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
