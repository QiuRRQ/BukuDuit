package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBooksTransactionRepository interface {
	BrowseByShop(userID string) (data []models.BooksTransaction, err error)

	Read(ID string) (data models.BooksTransaction, err error)

	Edit(body viewmodel.BooksTransactionVm) (res string, err error)

	Add(body viewmodel.BooksTransactionVm, userID string, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByShop(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByShop(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
