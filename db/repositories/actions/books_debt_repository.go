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

type BooksDebtRepository struct {
	DB *sql.DB
}

func NewBooksDebtModel(DB *sql.DB) contracts.IBooksDebtRepository {
	return BooksDebtRepository{DB: DB}
}

func (repository BooksDebtRepository) Browse(status string) (data []models.BooksDebt, err error) {

	statement := `select * from "books_debt" where "deleted_at" is null and "status" = $1`

	rows, err := repository.DB.Query(statement, status)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BooksDebt{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.CustomerID,
			&dataTemp.SubmissionDate,
			&dataTemp.BillDate,
			&dataTemp.DebtTotal,
			&dataTemp.CreditTotal,
			&dataTemp.Status,
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

func (repository BooksDebtRepository) BrowseByShop(shopID,status string) (data []models.BooksDebt, err error) {

	statement := `select bd.* from "books_debt" bd 
    inner join "user_customers" uc on uc."id"=bd."customer_id" 
    inner join "business_cards" bc on bc."id" = uc."business_id"
    where bd."deleted_at" is null and uc."business_id"=$1 and bd."status" = $2`

	rows, err := repository.DB.Query(statement, shopID,status)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.BooksDebt{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.CustomerID,
			&dataTemp.SubmissionDate,
			&dataTemp.BillDate,
			&dataTemp.DebtTotal,
			&dataTemp.CreditTotal,
			&dataTemp.Status,
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

func (repository BooksDebtRepository) BrowseByCustomer(customerID, status string) (data models.BooksDebt, err error) {
	if status == ""{
		statement := `select * from "books_debt" where "customer_id"=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, customerID).Scan(
			&data.ID,
			&data.CustomerID,
			&data.SubmissionDate,
			&data.BillDate,
			&data.DebtTotal,
			&data.CreditTotal,
			&data.Status,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
		)
	}else{
		statement := `select * from "books_debt" where "customer_id"=$1 and "deleted_at" is null and "status"=$2`
		err = repository.DB.QueryRow(statement, customerID, status).Scan(
			&data.ID,
			&data.CustomerID,
			&data.SubmissionDate,
			&data.BillDate,
			&data.DebtTotal,
			&data.CreditTotal,
			&data.Status,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
		)
	}
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository BooksDebtRepository) Read(ID string) (data models.BooksDebt, err error) {
	statement := `select * from "books_debt" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.CustomerID,
		&data.SubmissionDate,
		&data.BillDate,
		&data.DebtTotal,
		&data.CreditTotal,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository BooksDebtRepository) Edit(body viewmodel.BooksDebtVm, tx *sql.Tx) (res string, err error) {
	statement := `update "books_debt" set "bill_date"=$1, "debt_total"=$2, "credit_total"=$3, "status"=$4, "updated_at"=$5 where "id"=$6 returning "id"`
	err = tx.QueryRow(
		statement,
		str.EmptyString(body.BillDate),
		body.DebtTotal,
		body.CreditTotal,
		str.EmptyString(body.Status),
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksDebtRepository) Add(body viewmodel.BooksDebtVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "books_debt" ("customer_id","submission_date","bill_date","debt_total","credit_total","status","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7,$8) returning "id"`
	if tx != nil {
		err = tx.QueryRow(
			statement,
			str.EmptyString(body.CustomerID),
			body.SubmissionDate,
			str.EmptyString(body.BillDate),
			body.DebtTotal,
			body.CreditTotal,
			body.Status,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		).Scan(&res)
	} else {
		err = repository.DB.QueryRow(
			statement,
			str.EmptyString(body.CustomerID),
			body.SubmissionDate,
			str.EmptyString(body.BillDate),
			body.DebtTotal,
			body.CreditTotal,
			body.Status,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksDebtRepository) Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (err error) {
	statement := `update "books_debt" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	_,err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository BooksDebtRepository) DeleteByCustomer(customerID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "books_debt" set "updated_at"=$1, "deleted_at"=$2 where "customer_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), customerID)
	if err != nil {
		return err
	}

	return err
}

func (repository BooksDebtRepository) CountByCustomer(customerID, status string) (res int, err error) {
	statement := `select count("id") from "books_debt" where "customer_id"=$1 and "deleted_at" is null and "status"=$2`
	err = repository.DB.QueryRow(statement, customerID, status).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BooksDebtRepository) CountByPk(ID, status string) (res int, err error) {
	statement := `select count("id") from "books_debt" where "id"=$1 and "deleted_at" is null and "status"=$2`
	err = repository.DB.QueryRow(statement, ID, status).Scan(&res)
	if err != nil {
		fmt.Println(12)
		return res, err
	}

	return res, err
}

func (repository BooksDebtRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "books_debt" where ` + column + `=$1 and "deleted_at" is null and "status"=$2`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
