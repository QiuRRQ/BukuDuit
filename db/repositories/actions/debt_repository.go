package actions

import (
	"bukuduit-go/db/models"
	"bukuduit-go/db/repositories/contracts"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type DebtRepository struct{
	DB *sql.DB
}

func NewDebtModel(DB *sql.DB) contracts.IDebtsRepository{
	return DebtRepository{DB:DB}
}

func (repository DebtRepository) BrowseByCustomer(customerID, status string) (data []models.Debt, err error) {
	var rows *sql.Rows
	if status != ""{
		statement := `select * from "debts" where "customer_id"=$1 and "status"=$2 and "deleted_at" is null`
		rows,err = repository.DB.Query(statement,customerID,status)
		if err != nil {
			return data,err
		}
	}else{
		statement := `select * from "debts" where "customer_id"=$1 and "deleted_at" is null`
		rows,err = repository.DB.Query(statement,customerID)
		if err != nil {
			return data,err
		}
	}

	for rows.Next() {
		dataTemp := models.Debt{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.CustomerID,
			&dataTemp.SubmissionDate,
			&dataTemp.BillDate,
			&dataTemp.Total,
			&dataTemp.Status,
			&dataTemp.Type,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,err
		}

		data = append(data,dataTemp)
	}

	return data,err
}

func (repository DebtRepository) Read(ID string) (data models.Debt, err error) {
	statement := `select * from "debts" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,ID).Scan(
		&data.ID,
		&data.CustomerID,
		&data.SubmissionDate,
		&data.BillDate,
		&data.Total,
		&data.Status,
		&data.Type,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)
	if err != nil {
		return data,err
	}

	return data,nil
}

func (repository DebtRepository) Edit(ID, billDate, status string, debtType, total int32, tx *sql.Tx) (err error) {
	return err
}

func (repository DebtRepository) Add(body viewmodel.DebtVm, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repository DebtRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

