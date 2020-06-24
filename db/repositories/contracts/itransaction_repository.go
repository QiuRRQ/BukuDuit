package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type ITransactionRepository interface {
	BrowseByCustomer(customerID string) (data []models.Transactions, err error)

	Read(ID string) (data models.Transactions, err error)

	Edit(body viewmodel.TransactionVm) (res string, err error)

	Add(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByCustomer(customerID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID string) (res int, err error)

	DebtPayment(customerID, DebtType string, UserCustomerDebt, amount int) (CustomerDebt int)

	CountBy(column, value string) (res int, err error)
}
