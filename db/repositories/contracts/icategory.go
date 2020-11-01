package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type ICategory interface {
	Read() (data []models.Categories, err error)

	ReadByID(ID string) (data []models.Categories, err error)

	Edit(body viewmodel.CategoriesVM) (res string, err error)

	Add(body viewmodel.CategoriesVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
