package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type ITransactionRepository interface {
	BrowseByCustomer(customerID string) (data []models.Transactions, err error)

	BrowseByShop(shopID string) (data []models.Transactions, err error)

	TransactionBrowsByShop(shopID string) (data []models.Transactions, err error)

	Read(ID string) (data models.Transactions, err error)

	Edit(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error)

	Add(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByCustomer(customerID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)

	CountDistinctBy(column, ID string) (res int, err error)

	DebtReport(customerID, shopID, filter string) (data []models.Transactions, err error)
}
