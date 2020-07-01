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

type PaymentAccountRepository struct {
	DB *sql.DB
}

func NewPaymentAccountModel(DB *sql.DB) contracts.IPaymentAccountRepository {
	return PaymentAccountRepository{DB: DB}
}

func (repository PaymentAccountRepository) BrowseByShop(shopID string) (data []models.PaymentAccount, err error) {
	statement := `select * from "payment_accounts" where "shop_id"=$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement, shopID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.PaymentAccount{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.ShopID,
			&dataTemp.PaymentNumber,
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

func (repository PaymentAccountRepository) Read(ID string) (data models.PaymentAccount, err error) {
	statement := `select * from "payment_accounts" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.ShopID,
		&data.Name,
		&data.PaymentNumber,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository PaymentAccountRepository) Edit(body viewmodel.PaymentAccountVm) (res string, err error) {
	statement := `update "payment_accounts" set "shop_id"=$1, "name"=$2, "payment_number"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.ShopID,
		str.EmptyString(body.Name),
		str.EmptyString(body.PaymentNumber),
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository PaymentAccountRepository) Add(body viewmodel.PaymentAccountVm) (res string, err error) {
	statement := `insert into "payment_accounts" ("shopi_id","name","payment_number","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.ShopID,
		str.EmptyString(body.Name),
		str.EmptyString(body.PaymentNumber),
		datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
	).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository PaymentAccountRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "payment_accounts" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository PaymentAccountRepository) DeleteByShop(shopID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "payment_accounts" set "updated_at"=$1, "deleted_at"=$2 where "user_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), shopID)
	if err != nil {
		return err
	}

	return err
}

func (repository PaymentAccountRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "business_cards" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository PaymentAccountRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "business_cards" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
