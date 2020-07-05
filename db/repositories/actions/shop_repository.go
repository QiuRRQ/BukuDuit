package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/helpers/datetime"
	"bukuduit-go/helpers/str"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
	"fmt"
	"time"
)

type ShopRepository struct {
	DB *sql.DB
}

func NewShopModel(DB *sql.DB) contracts.IShopRepository {
	return ShopRepository{DB: DB}
}

func (repository ShopRepository) BrowseByUser(userID string) (data []models.Shop, err error) {
	statement := `select * from "shops" where "user_id"=$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement, userID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Shop{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.FullName,
			&dataTemp.BookName,
			&dataTemp.MobilePhone,
			&dataTemp.TagLine,
			&dataTemp.Address,
			&dataTemp.Email,
			&dataTemp.Avatar,
			&dataTemp.UserID,
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

func (repository ShopRepository) Read(ID string) (data models.Shop, err error) {
	statement := `select * from "shops" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.FullName,
		&data.BookName,
		&data.MobilePhone,
		&data.TagLine,
		&data.Address,
		&data.Email,
		&data.Avatar,
		&data.UserID,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository ShopRepository) Edit(body viewmodel.ShopVm) (res string, err error) {
	statement := `update "shops" set "full_name"=$1, "book_name"=$2, "mobile_phone"=$3, "tag_line"=$4, "address"=$5, "email"=$6, "avatar"=$7, "updated_at"=$8 where "id"=$9 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		str.EmptyString(body.FullName),
		body.BookName,
		str.EmptyString(body.MobilePhone),
		str.EmptyString(body.TagLine),
		str.EmptyString(body.Address),
		str.EmptyString(body.Email),
		str.EmptyString(body.Avatar),
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository ShopRepository) Add(body viewmodel.ShopVm, userID string, tx *sql.Tx) (res string, err error) {
	fmt.Println(body.MobilePhone)
	statement := `insert into "shops" ("full_name","book_name","mobile_phone","tag_line","address","email","avatar","user_id","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.FullName),
			body.BookName,
			str.EmptyString(body.MobilePhone),
			str.EmptyString(body.TagLine),
			str.EmptyString(body.Address),
			str.EmptyString(body.Email),
			str.EmptyString(body.Avatar),
			userID,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.FullName),
			body.BookName,
			str.EmptyString(body.MobilePhone),
			str.EmptyString(body.TagLine),
			str.EmptyString(body.Address),
			str.EmptyString(body.Email),
			str.EmptyString(body.Avatar),
			userID,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository ShopRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "shops" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository ShopRepository) DeleteByUser(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "shops" set "updated_at"=$1, "deleted_at"=$2 where "user_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), userID)
	if err != nil {
		return err
	}

	return err
}

func (repository ShopRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "shops" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository ShopRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "shops" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
