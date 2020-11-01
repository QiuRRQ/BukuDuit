package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IAuthors interface {
	Read() (data []models.Authors, err error)

	ReadByID(ID string) (data []models.Authors, err error)

	Edit(body viewmodel.AuthorsVM) (res string, err error)

	Add(body viewmodel.AuthorsVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
