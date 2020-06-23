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

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionModel(DB *sql.DB) contracts.ITransactionRepository {
	return TransactionRepository{DB: DB}
}

func (repository TransactionRepository) BrowseByCustomer(customerID string) (data []models.Transactions, err error) {
	statement := `select * from "transactions" where "customer_id"=$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement, customerID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Transactions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Amount,
			&dataTemp.Customer_Id,
			&dataTemp.Description,
			&dataTemp.Image,
			&dataTemp.Transaction_Date,
			&dataTemp.Type,
			&dataTemp.Created_at,
			&dataTemp.Update_at,
			&dataTemp.Deleted_at,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository TransactionRepository) Read(ID string) (data models.Transactions, err error) {
	statement := `select * from "transactions" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.Customer_Id,
		&data.Description,
		&data.Amount,
		&data.Created_at,
		&data.Deleted_at,
		&data.Image,
		&data.Transaction_Date,
		&data.Type,
		&data.Update_at,
		&data.Deleted_at,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository TransactionRepository) Edit(body viewmodel.TransactionVm) (res string, err error) {
	statement := `update "transactions" set "amount"=$1, "description"=$2, "image"=$3, "type"=$4, "transaction_date"=$5, "updated_at"=$6 where "id"=$7 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		str.EmptyString(body.Amount),
		str.EmptyString(body.Description),
		str.EmptyString(body.Image),
		str.EmptyString(body.Type),
		str.EmptyString(body.Transaction_Date),
		datetime.StrParseToTime(body.Update_at, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) Add(body viewmodel.TransactionVm, userID string, tx *sql.Tx) (res string, err error) {
	// fmt.Println(body.MobilePhone)
	statement := `insert into "transactions" ("customer_id","amount","description","image","type","transaction_date","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7,$8) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.Customer_Id),
			str.EmptyString(body.Amount),
			str.EmptyString(body.Description),
			str.EmptyString(body.Image),
			str.EmptyString(body.Type),
			datetime.StrParseToTime(body.Transaction_Date, time.RFC3339),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.Update_at, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.Customer_Id),
			str.EmptyString(body.Amount),
			str.EmptyString(body.Description),
			str.EmptyString(body.Image),
			str.EmptyString(body.Type),
			datetime.StrParseToTime(body.Transaction_Date, time.RFC3339),
			datetime.StrParseToTime(body.Created_at, time.RFC3339),
			datetime.StrParseToTime(body.Update_at, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "transactions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) DeleteByCustomer(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "updated_at"=$1, "deleted_at"=$2 where "user_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), userID)
	if err != nil {
		return err
	}

	return err
}

func (repository TransactionRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "transactions" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "transactions" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
