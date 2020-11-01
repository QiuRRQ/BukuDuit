package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IPublishers interface {
	Read() (data []models.Publishers, err error)

	ReadByID(ID string) (data []models.Publishers, err error)

	Edit(body viewmodel.PublishersVM) (res string, err error)

	Add(body viewmodel.PublishersVM, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
