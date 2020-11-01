package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBooksHasCategory interface {
	Read() (data []models.Books_has_category, err error)

	ReadByID(ID string) (data []models.Books_has_category, err error)

	Edit(body viewmodel.BooksHasCategoryVM) (res string, err error)

	Add(body viewmodel.BooksHasCategoryVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
