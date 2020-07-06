package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IDebtsRepository interface {
	BrowseByCustomer(customerID, status string) (data []models.Debt, err error)

	Read(ID string) (data models.Debt, err error)

	Edit(ID, billDate, status string, debtType, total int32, tx *sql.Tx) (err error)

	Add(body viewmodel.DebtVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error)
}
