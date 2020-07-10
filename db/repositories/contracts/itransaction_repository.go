package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type ITransactionRepository interface {
	BrowseByCustomer(customerID string) (data []models.Transactions, err error)

	BrowseByShop(shopID string) (data []models.Transactions, err error)

	TransactionBrowsByShop(shopID, filter string) (data []models.Transactions, err error)

	BrowseByBookDebtID(bookDebtID string, status int) (data []models.Transactions, err error)

	Read(ID string) (data models.Transactions, err error)

	Edit(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error)

	Add(body viewmodel.TransactionVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	DeleteByCustomer(customerID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)

	CountDistinctBy(column, ID string) (res int, err error)

	DebtReport(customerID, shopID, bookDebtID, filter string) (data []models.Transactions, err error)

	TransactionReport(shopID, filter string) (data []models.Transactions, err error)

	GroubByWeeksMonth(timeBy string) (data []models.TransByMonth, err error)

	FirstTransactionDate(shopId, filter string) (res string, err error)

	LastTransactionDate(shopId, filter string) (res string, err error)

	MakeWeeklySeries(startDate, endDate string) (res []models.Weekly, err error)
}
