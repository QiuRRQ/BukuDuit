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

type BooksTransactionRepository struct {
	DB *sql.DB
}

func NewBooksTransactionModel(DB *sql.DB) contracts.IBooksTransactionRepository {
	return BooksTransactionRepository{DB: DB}
}

func (repository BooksTransactionRepository) BrowseByShop(shopID string) (data []models.BooksTransaction, err error) {
	statement := `select * from "books_transactions" where "shop_id"=$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement, shopID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BooksTransaction{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.ShopID,
			&dataTemp.DebtTotal,
			&dataTemp.CreditTotal,
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

func (repository BooksTransactionRepository) Read(ID string) (data models.BooksTransaction, err error) {
	statement := `select * from "books_transactions" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.ShopID,
		&data.DebtTotal,
		&data.CreditTotal,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository BooksTransactionRepository) Edit(body viewmodel.BooksTransactionVm) (res string, err error) {
	statement := `update "books_transactions" set "debt_total"=$1, "credit_total"=$2,"updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		body.DebtTotal,
		body.CreditTotal,
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksTransactionRepository) Add(body viewmodel.BooksTransactionVm, userID string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "books_transactions" ("shop_id","debt_total","credit_total","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	if tx != nil {
		_, err = tx.Exec(
			statement,
			str.EmptyString(body.ShopID),
			body.DebtTotal,
			body.CreditTotal,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.ShopID),
			body.DebtTotal,
			body.CreditTotal,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksTransactionRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "books_transactions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksTransactionRepository) DeleteByShop(shopID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "books_transactions" set "updated_at"=$1, "deleted_at"=$2 where "shop_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), shopID)
	if err != nil {
		return err
	}

	return err
}

func (repository BooksTransactionRepository) CountByShop(shopID string) (res int, err error) {
	statement := `select count("id") from "books_transactions" where "shop_id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, shopID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksTransactionRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "books_transactions" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
