package contracts

import (
	"bukuduit-go/db/models"
	"bukuduit-go/usecase/viewmodel"
	"database/sql"
)

type IBusinessCardRepository interface {
	BrowseByUser(userID string) (data []models.BusinessCards, err error)

	Read(ID string) (data models.BusinessCards, err error)

	Edit(body viewmodel.BusinessCardVm, userID string) (res string, err error)

	Add(body viewmodel.BusinessCardVm, userID string, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByUser(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountByPk(ID string) (res int, err error)

	CountBy(column, value string) (res int, err error)
}
