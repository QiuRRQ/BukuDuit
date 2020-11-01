package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBooks interface {
	Read() (data []models.Books, err error)

	ReadByID(ID string) (data []models.Books, err error)

	Edit(body viewmodel.BooksVM) (res string, err error)

	EditStock(body viewmodel.BooksVM, tx *sql.Tx) (res string, err error)

	Add(body viewmodel.BooksVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
