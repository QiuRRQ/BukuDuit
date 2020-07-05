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

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionModel(DB *sql.DB) contracts.ITransactionRepository {
	return TransactionRepository{DB: DB}
}

func (repository TransactionRepository) DebtReport(customerID, shopID, bookDebtID, filter string) (data []models.Transactions, err error) {

	statement := `select t."id", uc."full_name", t."amount", t."reference_id", t."shop_id", t."description", t."image", t."transaction_date", t."type", t."created_at", t."updated_at", t."deleted_at" 
	from "transactions" t  join "user_customers" uc 
	on t."reference_id" = uc."id" 
	where t."books_debt_id" in (` + bookDebtID + `) and t."deleted_at" is null and t."customer_id" is null ` + filter

	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Transactions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.Amount,
			&dataTemp.ReferenceID,
			&dataTemp.IDShop,
			&dataTemp.Description,
			&dataTemp.Image,
			&dataTemp.TransactionDate,
			&dataTemp.Type,
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

//ini untuk list transaksi
func (repository TransactionRepository) TransactionBrowsByShop(shopID string) (data []models.Transactions, err error) {
	statement := `select t."id", uc."full_name", t."amount", t."reference_id", t."shop_id", t."description", t."image", t."transaction_date", t."type", t."created_at", t."updated_at", t."deleted_at" 
	from "transactions" t  join "user_customers" uc 
	on t."customer_id" = uc."id" 
	where t."shop_id" = $1 and t."deleted_at" is null and t."customer_id" is not null order by t."transaction_date" desc `

	rows, err := repository.DB.Query(statement, shopID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Transactions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Amount,
			&dataTemp.ReferenceID,
			&dataTemp.IDShop,
			&dataTemp.Description,
			&dataTemp.Image,
			&dataTemp.TransactionDate,
			&dataTemp.Type,
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

//ini untuk laporan utang
func (repository TransactionRepository) BrowseByShop(shopID string) (data []models.Transactions, err error) {
	statement := `select t."id", uc."full_name", t."amount", t."reference_id", t."shop_id", t."description", t."image", t."transaction_date", t."type", t."created_at", t."updated_at", t."deleted_at" 
	from "transactions" t  join "user_customers" uc 
	on t."reference_id" = uc."id" 
	where t."shop_id" = $1 and t."deleted_at" is null and t."customer_id" is null order by t."transaction_date" desc `

	rows, err := repository.DB.Query(statement, shopID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Transactions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.Amount,
			&dataTemp.ReferenceID,
			&dataTemp.IDShop,
			&dataTemp.Description,
			&dataTemp.Image,
			&dataTemp.TransactionDate,
			&dataTemp.Type,
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

//ini untuk detail utang
func (repository TransactionRepository) BrowseByCustomer(customerID string) (data []models.Transactions, err error) {
	statement := `select t."id", uc."full_name", t."amount", t."reference_id", t."shop_id", t."description", t."image", t."transaction_date", t."type", t."created_at", t."updated_at", t."deleted_at" 
	from "transactions" t  join "user_customers" uc 
	on t."reference_id" = uc."id" 
	where t."reference_id" = $1 and t."deleted_at" is null and t."customer_id" is null order by t."transaction_date" desc `
	rows, err := repository.DB.Query(statement, customerID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Transactions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.Amount,
			&dataTemp.ReferenceID,
			&dataTemp.IDShop,
			&dataTemp.Description,
			&dataTemp.Image,
			&dataTemp.TransactionDate,
			&dataTemp.Type,
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

func (repository TransactionRepository) Read(ID string) (data models.Transactions, err error) {
	statement := `select id, shop_id, reference_id, category_id, amount, description, image, type, transaction_date, created_at, 
	updated_at, deleted_at, status, customer_id, books_debt_id, books_transactions_id from "transactions" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(
		&data.ID,
		&data.IDShop,
		&data.ReferenceID,
		&data.CategoryID,
		&data.Amount,
		&data.Description,
		&data.Image,
		&data.Type,
		&data.TransactionDate,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Status,
		&data.CustomerID,
		&data.BooksDeptID,
		&data.BooksTransID,
	)
	if err != nil {
		return data, err
	}

	return data, err
}

func (repository TransactionRepository) Edit(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error) {
	fmt.Println(body.ID)
	statement := `update "transactions" set "amount"=$1, "description"=$2, "image"=$3, "type"=$4, "transaction_date"=$5, "updated_at"=$6 where "id"=$7 returning "id"`
	err = tx.QueryRow(
		statement,
		body.Amount,
		str.EmptyString(body.Description),
		str.EmptyString(body.Image),
		str.EmptyString(body.Type),
		str.EmptyString(body.TransactionDate),
		datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
		body.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) Add(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "transactions" ("reference_id", "shop_id", "amount","description","type","transaction_date","created_at","updated_at","customer_id", "books_debt_id", "books_transactions_id") 
	values($1,$2,$3,$4,$5,to_date($6, 'YYYY-MM-DD'),$7,$8, $9, $10, $11) returning "id"`

	if tx != nil {
		_, err = tx.Exec(
			statement,
			body.ReferenceID,
			body.ShopID,
			body.Amount,
			str.EmptyString(body.Description),
			str.EmptyString(body.Type),
			body.TransactionDate,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
			str.EmptyString(body.CustomerID),
			str.EmptyString(body.BooksDebtID),
			str.EmptyString(body.BooksTransID),
		)
	} else {
		err = repository.DB.QueryRow(
			statement,
			body.ReferenceID,
			body.ShopID,
			body.Amount,
			str.EmptyString(body.Description),
			str.EmptyString(body.Type),
			body.TransactionDate,
			datetime.StrParseToTime(body.CreatedAt, time.RFC3339),
			datetime.StrParseToTime(body.UpdatedAt, time.RFC3339),
			str.EmptyString(body.CustomerID),
			str.EmptyString(body.BooksDebtID),
			str.EmptyString(body.BooksTransID),
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning  "id"`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository TransactionRepository) DeleteByCustomer(referenceID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "updated_at"=$1, "deleted_at"=$2 where "reference_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), referenceID)
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

func (repository TransactionRepository) CountDistinctBy(column, ID string) (res int, err error) {
	fmt.Println(ID)
	fmt.Println(column)
	statement := `select count(distinct (` + column + `)) from "transactions" where ` + column + `=$1 and "deleted_at" is null`
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
