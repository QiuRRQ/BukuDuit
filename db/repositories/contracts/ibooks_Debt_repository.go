package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBooksDebtRepository interface {
	BrowseByCustomer(userID, status string) (data models.BooksDebt, err error)

	Browse(status string) (data []models.BooksDebt, err error)

	Read(ID string) (data models.BooksDebt, err error)

	Edit(body viewmodel.BooksDebtVm, tx *sql.Tx) (res string, err error)

	Add(body viewmodel.BooksDebtVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByCustomer(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID, status string) (res int, err error)

	CountBy(column, value string) (res int, err error)

	CountByCustomer(customerID, status string) (res int, err error)
}
